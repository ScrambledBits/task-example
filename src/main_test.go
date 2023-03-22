package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(hello)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "hello\n"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/headers", nil)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	req.Header.Set("Accept", "text/plain; charset=utf-8")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("X-Custom", "This is a custom header")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(headers)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "text/plain; charset=utf-8"
	if recorder.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			recorder.Header().Get("Content-Type"), expectedContentType)
	}
}
