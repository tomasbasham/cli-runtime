package flag_test

import (
	"strings"
	"testing"

	"github.com/tomasbasham/cli-runtime/flag"
)

func TestWordSepNormalizeFunc(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		flag string
		want string
	}{
		"hyphens are retained": {
			flag: "my-flag",
			want: "my-flag",
		},
		"underscores are replaced with hyphens": {
			flag: "my_flag",
			want: "my-flag",
		},
		"mixed separators": {
			flag: "my-flag_name",
			want: "my-flag-name",
		},
		"no separators": {
			flag: "myflag",
			want: "myflag",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			normaliser := flag.WordSepNormalizeFunc()

			got := string(normaliser(nil, tt.flag))
			if got != tt.want {
				t.Errorf("mismatch:\n  got:  %q\n  want: %q", got, tt.want)
			}
		})
	}
}

type mockWarner struct {
	buffer strings.Builder
}

func (m *mockWarner) Warn(message string) {
	m.buffer.WriteString(message)
}

func TestWarnWordSepNormalizerFunc(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		flag    string
		want    string
		wantMsg string
	}{
		"hyphens are retained": {
			flag: "my-flag",
			want: "my-flag",
		},
		"underscores are replaced with hyphens (with warning)": {
			flag:    "my_flag",
			want:    "my-flag",
			wantMsg: "flag name my_flag contains underscores, which are not allowed and have been replaced with hyphens",
		},
		"mixed separators (with warning)": {
			flag:    "my-flag_name",
			want:    "my-flag-name",
			wantMsg: "flag name my-flag_name contains underscores, which are not allowed and have been replaced with hyphens",
		},
		"no separators": {
			flag: "myflag",
			want: "myflag",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create a mock warner to capture warnings.
			mock := &mockWarner{}

			normaliser := flag.WarnWordSepNormalizeFunc(mock)

			got := string(normaliser(nil, tt.flag))
			if got != tt.want {
				t.Errorf("mismatch:\n  got:  %q\n  want: %q", got, tt.want)
			}
			if mock.buffer.String() != tt.wantMsg {
				t.Errorf("warning mismatch:\n  got:  %q\n  want: %q", mock.buffer.String(), tt.wantMsg)
			}
		})
	}
}
