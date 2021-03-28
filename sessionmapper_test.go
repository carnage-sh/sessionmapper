package sessionmapper

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionMapper(t *testing.T) {
	cfg := CreateConfig()
	cfg.Headers["me"] = "too"

	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	handler, err := New(ctx, next, cfg, "sessionmapper")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertString(t, "too", recorder.Header().Clone().Get("me"))
}

func assertString(t *testing.T, expected, value string) {
	t.Helper()

	if value != expected {
		t.Errorf("invalid value: %s, expected: %s", value, expected)
	}
}
