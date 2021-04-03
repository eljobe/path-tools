package cmd

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addToFront bool

func init() {
	addCmd.Flags().BoolVarP(&addToFront, "front", "f", false, "Prepend the path element to the front of the variable.")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add [-f] path_element",
	Short: "Add the given path element to the path variable if not already present.",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		parts := filepath.SplitList(os.Getenv(PathVar))
		seen := make(map[string]bool, len(parts))
		for _, part := range parts {
			seen[part] = true
		}
		pathel := args[0]
		if !seen[pathel] {
			if addToFront {
				parts = append([]string{pathel}, parts...)
			} else {
				parts = append(parts, pathel)
			}
		}
		fmt.Printf("%s=%s",PathVar,strings.Join(parts, string(os.PathListSeparator)))
		return nil
	},
}
