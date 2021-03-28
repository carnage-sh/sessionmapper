// a Traefik plugin that map session with additional header properties.
package sessionmapper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type client interface {
	Post(string, string, io.Reader) (*http.Response, error)
}

// Config the plugin configuration.
type Config struct {
	Headers []string `json:"headers,omitempty"`
	Server  string   `json:"server,omitempty"`
	Timeout int      `json:"timeout,omitempty"`
}

type Response struct {
	Upstream map[string]string `json:"upstream"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make([]string, 0),
		Server:  "http://localhost:7777/",
		Timeout: 200,
	}
}

// SessionMapper is the SessionMapper plugin.
type SessionMapper struct {
	headers []string
	client  client
	server  string
	next    http.Handler
}

// ServeHTTP implements the http.Handler
func (s *SessionMapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]string)
	for _, key := range s.headers {
		if v := r.Header.Get(key); v != "" {
			payload[key] = v
		}
	}
	b, _ := json.Marshal(&payload)
	fmt.Printf("server: %s", s.server)
	resp, err := s.client.Post(s.server, "application/json", bytes.NewReader(b))
	if err != nil {
		s.next.ServeHTTP(w, r)
		return
	}
	body := Response{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	err = json.Unmarshal(buf.Bytes(), &body)
	if err != nil {
		s.next.ServeHTTP(w, r)
		return
	}
	for k, v := range body.Upstream {
		r.Header.Add(k, v)
	}
	s.next.ServeHTTP(w, r)
}

// New created a new SessionMapper plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	client := &http.Client{
		// Timeout: time.Millisecond * time.Duration(config.Timeout),
	}
	return &SessionMapper{
		headers: config.Headers,
		client:  client,
		server:  config.Server,
		next:    next,
	}, nil
}
