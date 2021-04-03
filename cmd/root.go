package cmd

import (
	"github.com/spf13/cobra"
)

var PathVar string

var rootCmd = &cobra.Command{
	Use:   "path-tools [global flags]",
	Short: "A set of commands for managing your system's path variables.",
	Long: `The path-tools utility makes it easier to manage the path variables on your system.

Some example commands look like this:

path-tools add /usr/local/some-app/bin

path-tools ls

path-tools del /usr/local/some-app/bin

path-tools uniq
`,
	Version: "0.0.1",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&PathVar, "pathvar", "p", "PATH", "Name of the path variable.")
}
