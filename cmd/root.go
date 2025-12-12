package cmd

import (
	"fmt"
	"os"

	"gokit-init/internal/banner"
	"gokit-init/internal/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gokit-init",
	Short: "A Go project generator CLI tool",
	Long: `gokit-init is a CLI tool for generating boilerplate Go web application projects.
It supports multiple architectures, databases, and includes Docker support.`,
	Run: func(cmd *cobra.Command, args []string) {
		banner.Print()
		_ = cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gokit-init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gokit-init v%s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(newCmd)
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
