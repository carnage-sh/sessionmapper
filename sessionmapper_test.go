package sessionmapper

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockClient struct{}

func (m *MockClient) Post(string, string, io.Reader) (*http.Response, error) {
	r := bytes.NewBufferString(`{"upstream":{"they": "too"}}`)
	rc := io.NopCloser(r)
	return &http.Response{
		StatusCode: 200,
		Body:       rc,
	}, nil
}

func TestSessionMapper(t *testing.T) {
	cfg := CreateConfig()
	cfg.Headers = []string{"me"}

	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	handler := &SessionMapper{
		headers: []string{"they", "you"},
		client:  &MockClient{},
		server:  "server",
		next:    next,
	}
	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, req)
	assertString(t, "too", req.Header.Get("they"))
}

func assertString(t *testing.T, expected, value string) {
	t.Helper()

	if value != expected {
		t.Errorf(`invalid value: "%s" - expected: "%s"`, value, expected)
	}
}
