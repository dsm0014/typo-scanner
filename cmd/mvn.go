/*
Copyright © 2022 Daniel Morrison
*/
package cmd

import (
	"github.com/dsm0014/typo-scanner/scanner"
	"github.com/spf13/cobra"
)

var mvnCmd = &cobra.Command{
	Use:   "mvn [flags] [...package-names]",
	Short: "Scan Maven Repository for TypoSquatting on the input package-name(s) ",
	Long: `Scan the public Maven Repository at 'mvnrepository.com/artifact/' for TypoSquatting. 
Be sure to include the group name.

Example:
  typo-scanner mvn -dr io.cucumber/cucumber-java`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return scanner.TypoScanEntrypoint(scanner.ScanMvn, args, genFlags)
	},
}

func init() {
	rootCmd.AddCommand(mvnCmd)
}
