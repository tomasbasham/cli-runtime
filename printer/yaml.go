package printer

import (
	"bytes"
	"io"

	"github.com/goccy/go-yaml"
)

// YAMLPrinter is a [Printer] that outputs YAML formatted data.
type YAMLPrinter struct {
	// DocumentStart indicates whether to include the YAML document start marker
	// (---) to the first document in the output.
	DocumentStart bool
}

// Print writes the given value in YAML format to the given [io.Writer].
func (y *YAMLPrinter) Print(w io.Writer, v any) error {
	buf := &bytes.Buffer{}
	if y.DocumentStart {
		buf.WriteString("---\n")
	}

	if err := yaml.NewEncoder(buf).Encode(v); err != nil {
		return err
	}

	_, err := buf.WriteTo(w)
	return err
}
