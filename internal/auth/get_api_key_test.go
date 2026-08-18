package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		authHeader string
		setHeader  bool
		want       string
		wantErr    string
	}{
		"valid bearer token": {
			authHeader: "ApiKey abc123",
			setHeader:  true,
			want:       "abc123",
		},
		"no space to split": {
			authHeader: "abc123",
			setHeader:  true,
			wantErr:    "malformed authorization header", //malformed authorization header
		},
		"header missing": {
			setHeader: false,
			wantErr:   "no authorization header included",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			if tt.setHeader {
				header.Set("Authorization", tt.authHeader)
			}

			got, err := GetAPIKey(header)

			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
