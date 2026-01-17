package flag_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/pflag"

	"github.com/tomasbasham/cli-runtime/flag"
)

func TestPrinterFlags_AddFlags(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		format  string
		want    flag.Format
		wantErr bool
	}{
		"text format": {
			format:  "text",
			wantErr: false,
			want:    flag.FormatText,
		},
		"json format": {
			format:  "json",
			wantErr: false,
			want:    flag.FormatJSON,
		},
		"prettyjson format": {
			format:  "prettyjson",
			wantErr: false,
			want:    flag.FormatPrettyJSON,
		},
		"yaml format": {
			format:  "yaml",
			wantErr: false,
			want:    flag.FormatYAML,
		},
		"invalid format": {
			format:  "xml",
			wantErr: true,
		},
		"empty format": {
			format:  "",
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Set up FlagSet and add printer specific flags.
			flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
			pf := flag.NewPrinterFlags(
				flag.FormatTextFlag|flag.FormatJSONFlag|flag.FormatPrettyJSONFlag|flag.FormatYAMLFlag,
				flag.FormatText,
			)
			pf.AddFlags(flags)

			err := flags.Set("format", tt.format)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if !tt.wantErr {
				got := pf.Format
				if got != tt.want {
					t.Errorf("mismatch:\n  got:  %q\n  want: %q", got, tt.want)
				}
			}
		})
	}
}

type mockTextFormatter struct {
	text string
	err  error
}

func (m mockTextFormatter) FormatText() ([]byte, error) {
	if m.err != nil {
		return nil, m.err
	}
	return []byte(m.text), nil
}

func TestPrinterFlags_ToPrinter_TextFormatter(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input   any
		want    string
		wantErr bool
	}{
		"text formatter": {
			input:   mockTextFormatter{text: "formatted output"},
			want:    "formatted output",
			wantErr: false,
		},
		"text formatter with error": {
			input:   mockTextFormatter{err: errors.New("format error")},
			wantErr: true,
		},
		"non-text-formatter": {
			input:   "plain string",
			want:    "plain string\n",
			wantErr: false,
		},
		"non-text-formatter struct": {
			input:   map[string]string{"key": "value"},
			want:    "map[key:value]\n",
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Set up PrinterFlags with text format.
			pf := &flag.PrinterFlags{Format: flag.FormatText}

			p, err := pf.ToPrinter()
			if err != nil {
				t.Fatalf("failed to create printer: %v", err)
			}

			var buf bytes.Buffer
			err = p.Print(&buf, tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if !tt.wantErr {
				got := buf.String()
				if got != tt.want {
					t.Errorf("mismatch:\n  got:  %q\n  want: %q", got, tt.want)
				}
			}
		})
	}
}
