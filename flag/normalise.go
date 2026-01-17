package flag

import (
	"strings"

	"github.com/spf13/pflag"

	"github.com/tomasbasham/cli-runtime/printer"
)

type normalizer func(f *pflag.FlagSet, name string) pflag.NormalizedName

// discardWarner is a [printer.Warner] that does nothing.
type discardWarner struct{}

// Warn discards the warning message.
func (d discardWarner) Warn(message string) {}

// wordSepNormalizeFunc changes all underscores in given CLI flags to
// hyphens. No CLI application will accept underscores in flags, but instead of
// returning an error to the user we simply assume their intent and update their
// input.
func wordSepNormalizeFunc(w printer.Warner) normalizer {
	return func(f *pflag.FlagSet, name string) pflag.NormalizedName {
		if strings.Contains(name, "_") {
			w.Warn("flag name " + name + " contains underscores, which are not allowed and have been replaced with hyphens")
			return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
		}
		return pflag.NormalizedName(name)
	}
}

// WordSepNormalizeFunc returns a [normalizer] that replaces underscores in flag
// names with hyphens.
func WordSepNormalizeFunc() normalizer {
	return wordSepNormalizeFunc(discardWarner{})
}

// WarnWordSepNormalizeFunc returns a [normalizer] that replaces underscores in
// flag names with hyphens and prints a warning message.
func WarnWordSepNormalizeFunc(w printer.Warner) normalizer {
	return wordSepNormalizeFunc(w)
}
