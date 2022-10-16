/*
Copyright © 2022 Daniel Morrison
*/
package cmd

import (
	"github.com/dsm0014/typo-scanner/scanner"
	"github.com/dsm0014/typo-scanner/typo"
	"github.com/spf13/cobra"
)

var pypiCmd = &cobra.Command{
	Use:   "pypi [flags] [...package-names]",
	Short: "Scan PyPi for TypoSquatting on the input package-name(s) ",
	Long: `Scan the public PyPi registry 'pypi.org/project/' for TypoSquatting.

Example:
  typo-scanner pypi -dr cucumber`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, pkg := range args {
			typoList := typo.TypoGenerator(pkg, genFlags)
			_, err := scanner.ScanPypi(pkg, typoList)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pypiCmd)
}
