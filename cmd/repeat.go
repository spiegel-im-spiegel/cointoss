// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/cointoss/gen"
	"github.com/spiegel-im-spiegel/cointoss/repeat"
)

// repeatCmd represents the repeat command
var repeatCmd = &cobra.Command{
	Use:   "repeat",
	Short: "repeat coin tosses",
	Long:  "repeat coin tosses",
	Run: func(cmd *cobra.Command, args []string) {
		src, err := gen.NewRndSource(rngType, time.Now().UnixNano())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
			return
		}
		if err := repeat.Execute(repeat.NewContext(cmd.OutOrStdout(), os.Stderr, src, tossCount, repeatCount)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
		}
	},
}

func init() {
	RootCmd.AddCommand(repeatCmd)
	repeatCmd.PersistentFlags().Int64VarP(&repeatCount, "rcount", "c", 100, "Count of repetition")
}
