package cmd

import (
	"fmt"
	"path/filepath"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Lists the path elements in the path variable.",
	Aliases: []string{"ls"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		parts := filepath.SplitList(os.Getenv(PathVar))
		for _, part := range parts {
			fmt.Println(part)
		}
		return nil
	},
}
