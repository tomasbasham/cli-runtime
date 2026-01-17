package printer

import (
	"encoding/json"
	"io"
)

// JSONPrinter is a [Printer] that outputs JSON formatted data.
type JSONPrinter struct {
	// Indent indicates whether to pretty-print JSON output with indentation.
	Indent bool
}

// Print writes the given value in JSON format to the given [io.Writer].
func (f *JSONPrinter) Print(w io.Writer, v any) error {
	encoder := json.NewEncoder(w)
	if f.Indent {
		encoder.SetIndent("", "  ")
	}

	return encoder.Encode(v)
}
