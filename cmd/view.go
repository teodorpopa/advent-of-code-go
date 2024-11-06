package cmd

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/internal"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the puzzle description",
	Long: `View the puzzle description for the specified year and day.
If no year or date are specified, the latest available puzzle is displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		aocClient := internal.NewAocClient(Year, Day)
		fmt.Println(aocClient.ViewPuzzle())
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
