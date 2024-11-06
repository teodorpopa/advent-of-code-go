package cmd

import (
	"github.com/teodorpopa/advent-of-code-go/internal"

	"github.com/spf13/cobra"
)

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Scaffold a day to solve",
	Long: `Scaffold a day to solve.
This will download the puzzle description, 
input and also create the day files`,
	Run: func(cmd *cobra.Command, args []string) {
		aocClient := internal.NewAocClient(Year, Day)
		aocClient.GenerateDay()
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
}
