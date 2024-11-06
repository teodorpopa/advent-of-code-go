package cmd

import (
	"github.com/spf13/cobra"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"os"
)

var Year int
var Day int

var rootCmd = &cobra.Command{
	Use:   "advent-of-code-go",
	Short: "AoC client in Go",
	Long:  utils.DisplayTitle("Advent of Code in Go"),
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&Year, "year", "y", 0, "Year")
	rootCmd.PersistentFlags().IntVarP(&Day, "day", "d", 0, "Day")
}
