package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testCase struct {
	name            string
	path            string
	expectedCode    int
	expectedContent string
}

func setupAPI(t *testing.T) (url string, cleaner func()) {
	t.Helper()
	server := httptest.NewServer(newMux())
	url = server.URL
	cleaner = func() {
		server.Close()
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
			expectedContent: "404",
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
			case "text/plain; charset=utf-8":
				if !strings.Contains(body, tc.expectedContent) {
					t.Errorf("esperaba contenido %q, obtuve %q",
						tc.expectedContent, body)
				}
			default:
				t.Fatalf("Unsupported Content-Type: %q",
					resp.Header.Get("Content-Type"))
			}
		})
	}
}
