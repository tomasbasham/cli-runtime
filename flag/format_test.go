package flag_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/tomasbasham/cli-runtime/flag"
)

func TestFormatSet_Allows(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		flags  flag.FormatFlags
		format flag.Format
		want   bool
	}{
		"allow configured format (first in set)": {
			flags:  flag.FormatTextFlag | flag.FormatJSONFlag,
			format: flag.FormatText,
			want:   true,
		},
		"allow configured format (second in set)": {
			flags:  flag.FormatTextFlag | flag.FormatJSONFlag,
			format: flag.FormatJSON,
			want:   true,
		},
		"disallows unconfigured format": {
			flags:  flag.FormatTextFlag | flag.FormatJSONFlag,
			format: flag.FormatPrettyJSON,
			want:   false,
		},
		"disallows invalid format": {
			flags:  flag.FormatTextFlag | flag.FormatJSONFlag,
			format: "xml",
			want:   false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.flags.Allows(tt.format)
			if got != tt.want {
				t.Errorf("mismatch:\n  got:  %t\n  want: %t", got, tt.want)
			}
		})
	}
}

func TestFormatSet_AllowedFormats(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		flags flag.FormatFlags
		want  []flag.Format
	}{
		"multiple formats": {
			flags: flag.FormatTextFlag | flag.FormatJSONFlag,
			want:  []flag.Format{flag.FormatText, flag.FormatJSON},
		},
		"single format": {
			flags: flag.FormatYAMLFlag,
			want:  []flag.Format{flag.FormatYAML},
		},
		"no formats": {
			want: []flag.Format{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.flags.AllowedFormats()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestFormatSet_HelpString(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		flags flag.FormatFlags
		want  string
	}{
		"multiple formats": {
			flags: flag.FormatTextFlag | flag.FormatJSONFlag,
			want:  "text|json",
		},
		"single format": {
			flags: flag.FormatYAMLFlag,
			want:  "yaml",
		},
		"no formats": {
			want: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.flags.HelpString()
			if got != tt.want {
				t.Errorf("mismatch:\n  got:  %q\n  want: %q", got, tt.want)
			}
		})
	}
}
