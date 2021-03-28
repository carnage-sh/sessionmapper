package sessionmapper

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {

	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"upstream":{"me": "too"}}`)
			}))
	defer ts.Close()

	c := ts.Client()
	resp, err := c.Get(ts.URL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", resp.Status)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()
	assertString(t, `{"upstream":{"me": "too"}}`, s)
}

func TestSessionMapper(t *testing.T) {
	cfg := CreateConfig()
	cfg.Headers = []string{"me"}

	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"upstream":{"they": "too", "you": "too", "me": "too"}}`)
	}))
	defer ts.Close()
	c := ts.Client()

	handler := &SessionMapper{
		headers: []string{"they", "you"},
		client:  c,
		server:  ts.URL,
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
