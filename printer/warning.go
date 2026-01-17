package printer

import (
	"fmt"
	"io"
)

const (
	yellowColor = "\u001b[33;1m"
	resetColor  = "\u001b[0m"
)

// Warner is the interface implemented by a type that can print warnings.
type Warner interface {
	Warn(message string)
}

// WarningPrinter is an implementation of [Warner] that outputs warnings to the
// configured [io.Writer].
type WarningPrinter struct {
	out  io.Writer
	opts WarningPrinterOptions
}

// WarningPrinterOptions controls the behavior of a [WarningPrinter].
type WarningPrinterOptions struct {
	// Color indicates that warning output can include ANSI color codes.
	Color bool
}

// NewWarningPrinter returns an implementation of [WarningPrinter] that outputs
// warnings to the given [io.Writer].
// opts contains options controlling warning output
func NewWarningPrinter(out io.Writer, opts WarningPrinterOptions) *WarningPrinter {
	return &WarningPrinter{out, opts}
}

// Warn prints warnings to the configured [io.Writer].
func (w *WarningPrinter) Warn(message string) {
	if w.opts.Color {
		fmt.Fprintf(w.out, "%sWarning:%s %s\n", yellowColor, resetColor, message)
	} else {
		fmt.Fprintf(w.out, "Warning: %s\n", message)
	}
}
