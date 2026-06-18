package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/anasistemas/todo"
)

type testCase struct {
	name            string
	path            string
	expectedCode    int
	expectedContent string
	expItems        int
}

func setupAPI(t *testing.T) (url string, cleaner func()) {
	t.Helper()

	tf, err := os.CreateTemp("", "datafile*.json")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	server := httptest.NewServer(newMux(tf.Name()))

	for i := 1; i <= 3; i++ {
		var body bytes.Buffer

		type NewTask struct {
			Task string `json:"task"`
		}

		encoder := json.NewEncoder(&body)
		item := NewTask{Task: fmt.Sprintf("Task %d", i)}
		if err := encoder.Encode(item); err != nil {
			t.Fatal(err)
		}

		resp, err := http.Post(server.URL+"/todo", "application/json", &body)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("esperaba status %s al crear todo, obtuve %s",
				http.StatusText(http.StatusCreated),
				http.StatusText(resp.StatusCode))
		}
	}

	url = server.URL
	cleaner = func() {
		server.Close()
		os.Remove(tf.Name())
	}

	return url, cleaner
}

func TestGet(t *testing.T) {
	cases := []testCase{
		{
			name:            "Root",
			path:            "/",
			expectedCode:    http.StatusOK,
			expectedContent: "Hello World",
		},
		{
			name:            "NotFound",
			path:            "/blog",
			expectedCode:    http.StatusNotFound,
			expectedContent: "page not found",
		},
		{
			name:         "GetAll",
			path:         "/todo",
			expectedCode: http.StatusOK,
			expItems:     3,
		},
		{
			name:            "GetOne",
			path:            "/todo/1",
			expectedCode:    http.StatusOK,
			expItems:        1,
			expectedContent: "Task 1",
		},
	}

	url, cleaner := setupAPI(t)
	defer cleaner()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(url + tc.path)
			if err != nil {
				t.Fatalf("error haciendo la petición: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Errorf("esperaba status %s, obtuve %s",
					http.StatusText(tc.expectedCode),
					http.StatusText(resp.StatusCode))
			}

			buf := new(strings.Builder)
			io.Copy(buf, resp.Body)
			body := buf.String()

			switch resp.Header.Get("Content-Type") {
			case "text/plain", "text/plain; charset=utf-8":
				if !strings.Contains(body, tc.expectedContent) {
					t.Errorf("esperaba contenido %q, obtuve %q",
						tc.expectedContent, body)
				}

			case "application/json":
				var result struct {
					Results      todo.List `json:"results"`
					Date         time.Time `json:"date"`
					TotalResults int       `json:"total_results"`
				}

				if err := json.Unmarshal([]byte(body), &result); err != nil {
					t.Fatalf("error decodificando JSON: %v", err)
				}

				if result.TotalResults != tc.expItems {
					t.Errorf("esperaba %d items, obtuve %d",
						tc.expItems, result.TotalResults)
				}

				if tc.expectedContent != "" && len(result.Results) > 0 {
					firstTask := result.Results[0].Task
					if !strings.Contains(firstTask, tc.expectedContent) {
						t.Errorf("esperaba tarea %q, obtuve %q",
							tc.expectedContent, firstTask)
					}
				}

			default:
				t.Fatalf("Unsupported Content-Type: %q",
					resp.Header.Get("Content-Type"))
			}
		})
	}
}

func TestAdd(t *testing.T) {
	url, cleaner := setupAPI(t)
	defer cleaner()

	taskName := "Task 4"

	t.Run("Add", func(t *testing.T) {
		var body bytes.Buffer

		type NewTask struct {
			Task string `json:"task"`
		}

		item := NewTask{
			Task: taskName,
		}

		encoder := json.NewEncoder(&body)

		if err := encoder.Encode(item); err != nil {
			t.Fatal(err)
		}

		resp, err := http.Post(
			url+"/todo",
			"application/json",
			&body,
		)

		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Errorf(
				"esperaba status %s, obtuve %s",
				http.StatusText(http.StatusCreated),
				http.StatusText(resp.StatusCode),
			)
		}
	})

	t.Run("CheckAdd", func(t *testing.T) {
		resp, err := http.Get(url + "/todo/4")

		if err != nil {
			t.Fatal(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf(
				"esperaba status %s, obtuve %s",
				http.StatusText(http.StatusOK),
				http.StatusText(resp.StatusCode),
			)
		}

		var result struct {
			Results      todo.List `json:"results"`
			Date         time.Time `json:"date"`
			TotalResults int       `json:"total_results"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}

		if len(result.Results) == 0 {
			t.Fatal("no se encontraron resultados")
		}

		if result.Results[0].Task != taskName {
			t.Errorf(
				"esperaba tarea %q, obtuve %q",
				taskName,
				result.Results[0].Task,
			)
		}
	})
}

func TestDelete(t *testing.T) {
	url, cleaner := setupAPI(t)
	defer cleaner()

	t.Run("Delete", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, url+"/todo/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			t.Errorf(
				"esperaba status %s, obtuve %s",
				http.StatusText(http.StatusNoContent),
				http.StatusText(resp.StatusCode),
			)
		}
	})

	t.Run("CheckDelete", func(t *testing.T) {
		resp, err := http.Get(url + "/todo")
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf(
				"esperaba status %s, obtuve %s",
				http.StatusText(http.StatusOK),
				http.StatusText(resp.StatusCode),
			)
		}

		var result struct {
			Results      todo.List `json:"results"`
			Date         time.Time `json:"date"`
			TotalResults int       `json:"total_results"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}

		if result.TotalResults != 2 {
			t.Errorf(
				"esperaba 2 items, obtuve %d",
				result.TotalResults,
			)
		}

		if result.Results[0].Task != "Task 2" {
			t.Errorf(
				"esperaba Task 2, obtuve %q",
				result.Results[0].Task,
			)
		}
	})
}
