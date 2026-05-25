package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupAPI(t *testing.T) (url string, cleaner func()) {
	t.Helper()
	server := httptest.NewServer(newMux())
	url = server.URL
	cleaner = func() {
		server.Close()
	}
	return url, cleaner
}

func TestGetRoot(t *testing.T) {
	path := "/"
	expectedCode := http.StatusOK
	expectedContent := "Hello World"

	url, cleaner := setupAPI(t)
	defer cleaner()

	resp, err := http.Get(url + path)
	if err != nil {
		t.Fatalf("error haciendo la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedCode {
		t.Errorf("esperaba status %s, obtuve %s",
			http.StatusText(expectedCode),
			http.StatusText(resp.StatusCode))
	}

	buf := new(strings.Builder)
	io.Copy(buf, resp.Body)
	body := buf.String()

	if !strings.Contains(body, expectedContent) {
		t.Errorf("esperaba contenido %q, obtuve %q", expectedContent, body)
	}
}

func TestGetNotFound(t *testing.T) {
	path := "/blog"
	expectedCode := http.StatusNotFound
	expectedContent := "404"

	url, cleaner := setupAPI(t)
	defer cleaner()

	resp, err := http.Get(url + path)
	if err != nil {
		t.Fatalf("error haciendo la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedCode {
		t.Errorf("esperaba status %s, obtuve %s",
			http.StatusText(expectedCode),
			http.StatusText(resp.StatusCode))
	}

	buf := new(strings.Builder)
	io.Copy(buf, resp.Body)
	body := buf.String()

	if !strings.Contains(body, expectedContent) {
		t.Errorf("esperaba contenido %q, obtuve %q", expectedContent, body)
	}
}
