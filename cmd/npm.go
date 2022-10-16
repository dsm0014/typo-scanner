/*
Copyright © 2022 Daniel Morrison
*/
package cmd

import (
	"github.com/dsm0014/typo-scanner/scanner"
	"github.com/dsm0014/typo-scanner/typo"
	"github.com/spf13/cobra"
)

var npmCmd = &cobra.Command{
	Use:   "npm [flags] [...package-names]",
	Short: "Scan NPM for TypoSquatting on the input package-name(s) ",
	Long: `Scan the public NPM registry at 'npmjs.com/package/' for TypoSquatting.

Example:
  typo-scanner npm -d react`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, pkg := range args {
			typoList := typo.TypoGenerator(pkg, genFlags)
			_, err := scanner.ScanNpm(pkg, typoList)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(npmCmd)
}
