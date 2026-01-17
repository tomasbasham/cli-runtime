package cliruntime

import "github.com/spf13/cobra"

// CommandGroup is a collection of commands that are logically grouped together.
// The [CommandGroup.Title] field is used to identify the group.
type CommandGroup struct {
	Title    string
	Commands []*cobra.Command
}

// CommandGroups is a collection of [CommandGroup] instances. This type is used
// to register multiple command groups with a [cobra.Command].
type CommandGroups []CommandGroup

// Add registers each [CommandGroup] instance with the provided [cobra.Command].
func (g CommandGroups) Add(cmd *cobra.Command) {
	for _, group := range g {
		registerCommandGroup(cmd, group)
	}
}

func registerCommandGroup(cmd *cobra.Command, cg CommandGroup) {
	if len(cg.Title) == 0 {
		panic("CommandGroup requires a name")
	}

	group := &cobra.Group{
		ID:    cg.Title,
		Title: cg.Title,
	}

	for _, cc := range cg.Commands {
		cc.GroupID = cg.Title
		cmd.AddCommand(cc)
	}

	cmd.AddGroup(group)
}
