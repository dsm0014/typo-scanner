/*
Copyright © 2022 Daniel Morrison
*/
package cmd

import (
	"github.com/dsm0014/typo-scanner/typo"
	"os"

	"github.com/spf13/cobra"
)

var (
	genFlags typo.GeneratorFlags
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "typo-scanner",
	Short: "A scanner tool for detecting TypoSquatting",
	Long: `The CLI checks against common public registries for packages which may be
attempting to impersonate your package and wreak havoc via TypoSquatting.

The typo-scanner CLI is a lightweight tool for discovering if your packages
are being subject to this form of Software Supply Chain attack. The scanner 
generates a multitude of types of typos and verifies whether they exist or not.

Examples:
  typo-scanner npm -dr react
  typo-scanner pypi -d fastapi -x faastapi
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// GLOBAL FLAGS
	// -- Typo Variants --
	rootCmd.PersistentFlags().BoolVarP(&genFlags.Typo.ExtraKey, "extra-key", "e", false, "Check for typos with an additional character")
	rootCmd.PersistentFlags().BoolVarP(&genFlags.Typo.Skip, "skip", "s", false, "Check for typos with skipped characters")
	rootCmd.PersistentFlags().BoolVarP(&genFlags.Typo.Double, "double", "d", false, "Check for typos with doubled characters")
	rootCmd.PersistentFlags().BoolVarP(&genFlags.Typo.Reverse, "reverse", "r", false, "Check for typos with reversed characters")
	rootCmd.PersistentFlags().BoolVarP(&genFlags.Typo.Vowel, "vowel", "v", false, "Check for typos with incorrect vowels")
	rootCmd.PersistentFlags().BoolVarP(&genFlags.Typo.Key, "key", "k", false, "Check for typos with any incorrect characters")

	//
	rootCmd.PersistentFlags().StringSliceVarP(&genFlags.Excluded, "excluded", "x", []string{}, "Array of typos to exclude from scans (ex: -x faastapi,fasttapi)")

	// -- Logs --
	rootCmd.PersistentFlags().BoolVarP(&genFlags.SuppressLogs, "quiet", "q", false, "suppress all logs")

}
