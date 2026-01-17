package flag

import "strings"

// Format represents the output format.
type Format string

const (
	FormatText       Format = "text"
	FormatJSON       Format = "json"
	FormatPrettyJSON Format = "prettyjson"
	FormatYAML       Format = "yaml"
)

// FormatFlags represents a set of allowed output formats.
type FormatFlags uint8

const (
	FormatTextFlag FormatFlags = 1 << iota
	FormatJSONFlag
	FormatPrettyJSONFlag
	FormatYAMLFlag
)

var formatTable = []struct {
	flag   FormatFlags
	format Format
}{
	{FormatTextFlag, FormatText},
	{FormatJSONFlag, FormatJSON},
	{FormatPrettyJSONFlag, FormatPrettyJSON},
	{FormatYAMLFlag, FormatYAML},
}

// Allows checks if the [FormatFlags] allows the given Format.
func (fs FormatFlags) Allows(f Format) bool {
	for _, e := range formatTable {
		if e.format == f {
			return fs&e.flag != 0
		}
	}
	return false
}

// AllowedFormats returns a slice of allowed formats in the [FormatFlags].
func (fs FormatFlags) AllowedFormats() []Format {
	out := []Format{}
	for _, e := range formatTable {
		if fs&e.flag != 0 {
			out = append(out, e.format)
		}
	}
	return out
}

// HelpString returns a string representation of allowed formats for help
// messages.
func (fs FormatFlags) HelpString() string {
	formats := fs.AllowedFormats()
	strs := make([]string, len(formats))
	for i, f := range formats {
		strs[i] = string(f)
	}
	return strings.Join(strs, "|")
}
