package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name   string
		header []string
		want   string
		err    error
	}{
		{name: "good header", header: []string{"Authorization", "ApiKey key"}, want: "key", err: nil},
		{name: "good header1", header: []string{"Authorization", "ApiKey key1"}, want: "key1", err: nil},
		{name: "Bad Value header", header: []string{"Authorization", "key"}, want: "", err: errors.New("malformed authorization header")},
		{name: "Bad Key header", header: []string{"Author", "ApiKey key"}, want: "", err: ErrNoAuthHeaderIncluded},
		{name: "Bad Malformed Value header", header: []string{"Authorization", "ApiKeykey"}, want: "", err: errors.New("malformed authorization header")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := http.Header{}
			header.Set(tt.header[0], tt.header[1])

			got, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tt.err.Error()) {
					return
				}
				t.Fatalf("expected: %s, got: %s", err, tt.err)
			}
			if got != tt.want {
				t.Fatalf("expected: %s, got: %s, err: %s", tt.want, got, err)
			}

		})
	}
}
