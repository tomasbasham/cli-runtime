package templates

import (
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

const Indentation = "  "

// LongDesc normalizes a command's long description to follow the conventions.
func LongDesc(s string) string {
	return RenderMarkdownForCobra(heredoc.Doc(s))
}

// Usage normalizes a command's usage to follow the conventions.
func Usage(s string) string {
	return RenderMarkdownForCobra(heredoc.Doc(s))
}

// Examples normalizes a command's examples to follow the conventions.
func Examples(s string) string {
	if len(s) == 0 {
		return s
	}

	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i, l := range lines {
		lines[i] = Indentation + strings.TrimSpace(l)
	}
	return strings.Join(lines, "\n")
}

// Normalize perform all required normalizations on a given [cobra.Command].
func Normalize(cmd *cobra.Command) *cobra.Command {
	if len(cmd.Long) > 0 {
		cmd.Long = LongDesc(cmd.Long)
	}
	if len(cmd.Example) > 0 {
		cmd.Example = Examples(cmd.Example)
	}
	return cmd
}

// NormalizeAll perform all required normalizations in the entire
// [cobra.Command] tree.
func NormalizeAll(cmd *cobra.Command) *cobra.Command {
	if cmd.HasSubCommands() {
		for _, subCmd := range cmd.Commands() {
			NormalizeAll(subCmd)
		}
	}
	Normalize(cmd)
	return cmd
}
