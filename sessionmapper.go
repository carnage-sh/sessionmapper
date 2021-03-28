// a Traefik plugin that map session with additional header properties.
package sessionmapper

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

// SessionMapper is the SessionMapper plugin.
type SessionMapper struct {
	next    http.Handler
	headers map[string]string
	name    string
}

func (s *SessionMapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("me", "too")
	s.next.ServeHTTP(w, r)
}

// New created a new SessionMapper plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Headers) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &SessionMapper{
		headers: config.Headers,
		next:    next,
		name:    name,
	}, nil
}
