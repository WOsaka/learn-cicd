package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		err   string // Change to string for error message comparison
	}{
		"valid header": {
			input: http.Header{"Authorization": []string{"ApiKey abc123"}},
			want:  "abc123",
			err:   "",
		},
		"no header": {
			input: http.Header{},
			want:  "",
			err:   "no authorization header included", // Compare error message directly
		},
		"malformed header": {
			input: http.Header{"Authorization": []string{"Bearer abc123"}},
			want:  "",
			err:   "malformed authorization header", // Compare error message directly
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if got != tc.want || (err != nil && err.Error() != tc.err) {
				t.Fatalf("got: %q, want: %q, got error: %v, want error: %q", got, tc.want, err, tc.err)
			}
		})
	}
}