package printer_test

import (
	"bytes"
	"testing"

	"github.com/tomasbasham/cli-runtime/printer"
)

func TestWarningPrinter_Print(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		msg   string
		want  string
		color bool
	}{
		"message with color": {
			msg:   "this is a warning",
			want:  "\u001b[33;1mWarning:\u001b[0m this is a warning\n",
			color: true,
		},
		"message without color": {
			msg:   "this is a warning",
			want:  "Warning: this is a warning\n",
			color: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			w := printer.NewWarningPrinter(&buf, printer.WarningPrinterOptions{Color: tt.color})
			w.Warn(tt.msg)

			got := buf.String()
			if got != tt.want {
				t.Errorf("message mismatch:\n  got:  %q\n  want: %q", got, tt.want)
			}
		})
	}
}
