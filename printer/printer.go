package printer

import "io"

// TextFormatter is the interface implemented by a type that can marshal itself
// to text.
//
// [TextFormatter.FormatText] should return the encoded text or an error if
// the formatting operation fails.
type TextFormatter interface {
	FormatText() ([]byte, error)
}

// Printer operates on a value and writes the output to the given [io.Writer].
//
// [Printer.Print] should return an error if the printing operation fails.
type Printer interface {
	Print(io.Writer, any) error
}

// PrinterFunc type is an adapter to allow the use of ordinary functions as a
// printer. If f is a function with the appropriate signature, PrinterFunc(f) is
// a [Printer] that calls f.
type PrinterFunc func(io.Writer, any) error

// Print calls the [PrinterFunc] with the given [io.Writer] and value.
func (f PrinterFunc) Print(w io.Writer, v any) error {
	return f(w, v)
}
