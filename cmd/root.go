/*
Copyright © 2022 Daniel Morrison
*/
package cmd

import (
	"log"
	"os"
	"github.com/dsm0014/typo-scanner/typo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	typoFlags typo.TypoFlags
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
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	if typoFlags == typo.NewTypoFlags() {
		log.Println("Error: At least one typo flag must be specified")
		rootCmd.Help()
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.typo-scanner.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&typoFlags.ExtraKey, "extra-key", "e", false, "Check for typos with an additional character")
	rootCmd.PersistentFlags().BoolVarP(&typoFlags.Skip, "skip", "s", false, "Check for typos with skipped characters")
	rootCmd.PersistentFlags().BoolVarP(&typoFlags.Double, "double", "d", false, "Check for typos with doubled characters")
	rootCmd.PersistentFlags().BoolVarP(&typoFlags.Reverse, "reverse", "r", false, "Check for typos with reversed characters")
	rootCmd.PersistentFlags().BoolVarP(&typoFlags.Vowel, "vowel", "v", false, "Check for typos with incorrect vowels")
	rootCmd.PersistentFlags().BoolVarP(&typoFlags.Key, "key", "k", false, "Check for typos with any incorrect characters")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".typo-scanner" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".typo-scanner")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Fatal(err)

	}
}
