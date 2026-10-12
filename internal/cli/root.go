package cli

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/yanglichao/codekb/internal/exitcode"
)

// NewRootCommand builds the CodeKB command tree.
func NewRootCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "codekb",
		Short:   "Build and query a local code knowledge base",
		Version: version,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	return cmd
}

// Run executes the CLI and converts failures into process exit codes.
func Run(args []string, stdout, stderr io.Writer, version string) int {
	cmd := NewRootCommand(version)
	cmd.SetArgs(args)
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return exitcode.Failure
	}
	return exitcode.Success
}
