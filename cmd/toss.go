/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/cointoss/gen"
	"github.com/spiegel-im-spiegel/cointoss/toss"
)

// tossCmd represents the toss command
var tossCmd = &cobra.Command{
	Use:   "toss",
	Short: "coin toss",
	Long:  "coin toss",
	Run: func(cmd *cobra.Command, args []string) {
		src, err := gen.NewRndSource(rngType, time.Now().UnixNano())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
			return
		}
		if _, err := toss.Execute(toss.NewContext(cmd.OutOrStdout(), os.Stderr, src, tossCount)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
		}
	},
}

func init() {
	RootCmd.AddCommand(tossCmd)
}
