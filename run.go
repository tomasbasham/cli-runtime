package cliruntime

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tomasbasham/cli-runtime/flag"
)

// Run executes the provided [cobra.Command] and returns the exit code. Any
// errors that occur during command execution are printed to stderr.
func Run(cmd *cobra.Command) int {
	if err := run(cmd); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}
	return 0
}

// RunNoErrOutput executes the provided [cobra.Command] and returns any errors
// that occur during command execution.
func RunNoErrOutput(cmd *cobra.Command) error {
	return run(cmd)
}

func run(cmd *cobra.Command) error {
	cmd.SetGlobalNormalizationFunc(flag.WordSepNormalizeFunc())

	// When error printing is enabled for the Cobra command, a flag parse error
	// gets printed first, then optionally the often long usage text. This is very
	// unreadable in a console because the last few lines that will be visible on
	// screen don't include the error.
	//
	// Instead print the usage text, then the error. We implement this
	// consistently for all commands here. However, we don't want to print the
	// usage text when command execution fails for other reasons than parsing. We
	// detect this via the FlagParseError callback.
	//
	// Some commands, like kubectl, already deal with this themselves. We don't
	// change the behavior for those.
	if !cmd.SilenceUsage {
		cmd.SilenceUsage = true
		cmd.SetFlagErrorFunc(func(c *cobra.Command, err error) error {
			// Re-enable usage printing.
			c.SilenceUsage = false
			return err
		})
	}

	// In all cases error printing is done below.
	cmd.SilenceErrors = true

	return cmd.Execute()
}
