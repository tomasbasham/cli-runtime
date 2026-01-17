package printer

import "io"

// Discard is an instance of a [Printer] that discards all messages.
var Discard = discard{}

// Discard is a [Printer] that discards all messages.
type discard struct{}

// Print discards the value and returns nil.
func (d discard) Print(w io.Writer, v any) error {
	return nil
}
