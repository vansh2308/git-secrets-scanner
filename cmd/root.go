package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vansh2308/git-secrets-scanner.git/cmd/scm"
)

var rootCmd = &cobra.Command{
	Use:   "git-secret-scanner",
	Short: "Scan for secrets in your organization combining Gileaks and TruffleHog.",
}

func init() {
	// disable help command, prefer flags
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// register commands
	rootCmd.AddCommand(versionCmd)
	scm.AddScmCommands(rootCmd)

	// help flag
	rootCmd.Flags().BoolP("help", "h", false, "Help for git-secret-scanner")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
