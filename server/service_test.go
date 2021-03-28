package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceHandlerWithGet(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("me", "too")
	ServiceHandler(recorder, req)
	assertString(t, `{"upstream":{"they":"too"}}`, recorder.Body.String())
}

func TestServiceHandlerWithPost(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()

	data := bytes.NewReader([]byte(`{"me": "too"}`))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost", data)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("me", "too")
	ServiceHandler(recorder, req)
	assertString(t, `{"upstream":{"they":"too","you":"too"}}`, recorder.Body.String())
}

func assertString(t *testing.T, expected, value string) {
	t.Helper()

	if value != expected {
		t.Errorf("invalid value: %s, expected: %s", value, expected)
	}
}
