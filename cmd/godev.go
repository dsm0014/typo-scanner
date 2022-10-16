/*
Copyright © 2022 Daniel Morrison
*/
package cmd

import (
	"github.com/dsm0014/typo-scanner/scanner"
	"github.com/spf13/cobra"
)

var godevCmd = &cobra.Command{
	Use:   "godev [flags] [...package-names]",
	Short: "Scan go.dev for TypoSquatting on the input package-name(s) ",
	Long: `Scan the public go module registry at 'pkg.go.dev/' for TypoSquatting. 
Be sure to include the group name.

Example:
  typo-scanner go -dr github.com/fsnotify/fsnotify`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, pkg := range args {
			_, err := scanner.ScanGo(pkg, genFlags)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(godevCmd)
}
