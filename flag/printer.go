package flag

import (
	"fmt"
	"io"

	"github.com/spf13/pflag"

	"github.com/tomasbasham/cli-runtime/printer"
)

// PrinterFlags holds printer flags for CLI output formatting.
type PrinterFlags struct {
	Format          Format
	AcceptedFormats FormatFlags
}

// NewPrinterFlags creates a new [PrinterFlags] with the given accepted formats
// and default format.
func NewPrinterFlags(accepted FormatFlags, defaultFormat Format) *PrinterFlags {
	if !accepted.Allows(defaultFormat) {
		panic("default format not allowed by accepted format set")
	}

	return &PrinterFlags{
		AcceptedFormats: accepted,
		Format:          defaultFormat,
	}
}

// AddFlags adds flags to the given [pglag.FlagSet].
func (f *PrinterFlags) AddFlags(flags *pflag.FlagSet) {
	help := fmt.Sprintf("Output format (%s)", f.AcceptedFormats.HelpString())
	flags.FuncP("format", "f", help, func(s string) error {
		format := Format(s)

		if !f.AcceptedFormats.Allows(format) {
			return fmt.Errorf("invalid format %q (must be one of %s)", s, f.AcceptedFormats.HelpString())
		}

		f.Format = format
		return nil
	})
}

// ToPrinter converts the [PrinterFlags] to a [printer.Printer].
func (f *PrinterFlags) ToPrinter() (printer.Printer, error) {
	switch f.Format {
	case FormatYAML:
		return &printer.YAMLPrinter{DocumentStart: true}, nil
	case FormatJSON:
		return &printer.JSONPrinter{}, nil
	case FormatPrettyJSON:
		return &printer.JSONPrinter{Indent: true}, nil
	case FormatText:
		return printer.PrinterFunc(func(w io.Writer, a any) error {
			tf, ok := a.(printer.TextFormatter)
			if !ok {
				_, err := fmt.Fprintln(w, a)
				return err
			}

			b, err := tf.FormatText()
			if err != nil {
				return err
			}

			_, err = w.Write(b)
			return err
		}), nil
	default:
		return nil, fmt.Errorf("unknown format: %s", f.Format)
	}
}
