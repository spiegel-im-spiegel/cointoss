/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// Flags
var (
	ExitCode    int
	rngType     string
	tossCount   int64
	repeatCount int64
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cointoss",
	Short: "simuration of coin toss",
	Long:  "simuration of coin toss.",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		ExitCode = 1
	}
}

func init() {
	ExitCode = 0
	RootCmd.PersistentFlags().Int64VarP(&tossCount, "toss", "t", 10, "Count of tosses")
	RootCmd.PersistentFlags().StringVarP(&rngType, "rsource", "r", "GO", "Source of RNG (GO/MT)")
}
