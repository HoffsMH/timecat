package main

import (
	"github.com/spf13/cobra"
	"github.com/hoffsmh/timecat"
	"fmt"
	"os"
)

var days int
var weeks int
var months int
var dir string

var rootCmd = &cobra.Command{
	Use: "stuff",
	Short: "cats together files whose names form a specified time window relative to now",
	Long: `cats together files whose names form a specified time window relative to now

For example...

Given this:
	
./somedir
├── 2021-08-08T14:13:45+0000-sometext.md
├── 2021-10-01-blah.md
├── 2021-10-08T14:13:45+0000-someothertext.md
├── 2021-10-11-sometext.md
├── 2021-10-12T14:13:45-sometext.md
└── 2021-9-29T14:13:45+0000-sometext.md

Try this:
	timecat  somedir --months 1000`,
	Run: func(cmd *cobra.Command,  args []string) {
		dir = "."
		if len(args) > 0 {
			dir = args[0]
		}
		fmt.Println(timecat.Cat(args[0], &timecat.TimeRange{months,weeks,days}))
	},
}

func main() {
 if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  rootCmd.Flags().IntVarP(&days, "days", "d", 1, "the amount of days to look back.(1)")
  rootCmd.Flags().IntVarP(&weeks, "weeks", "w", 0, "the amount of weeks to look back.(0)")
  rootCmd.Flags().IntVarP(&months, "months", "m", 0, "the amount of months to look back.(0)")
}