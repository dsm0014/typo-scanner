/*
Copyright © 2022 Daniel Morrison
*/
package cmd

import (
	"github.com/dsm0014/typo-scanner/scanner"
	"github.com/spf13/cobra"
)

var rubyCmd = &cobra.Command{
	Use:   "ruby [flags] [...package-names]",
	Short: "Scan Ruby Gems for TypoSquatting on the input package-name(s) ",
	Long: `Scan the public Ruby Gem registry 'rubygems.org/gems/' for TypoSquatting.

Example:
  typo-scanner ruby -dr cucumber`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return scanner.TypoScanEntrypoint(scanner.ScanRuby, args, genFlags)
	},
}

func init() {
	rootCmd.AddCommand(rubyCmd)
}
