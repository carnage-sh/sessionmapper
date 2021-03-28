package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDisplayHandler(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("me", "too")
	DisplayHandler(recorder, req)
	assertString(t, "Headers\nMe: too\n", recorder.Body.String())
}

func assertString(t *testing.T, expected, value string) {
	t.Helper()

	if value != expected {
		t.Errorf("invalid value: %s, expected: %s", value, expected)
	}
}
