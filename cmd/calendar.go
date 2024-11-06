package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/teodorpopa/advent-of-code-go/internal"
)

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Display the AoC calendar",
	Long: `Display the AoC calendar for the specified year.
If no year is specified, the current year will be used.
`,
	Run: func(cmd *cobra.Command, args []string) {
		aocClient := internal.NewAocClient(Year, Day)
		fmt.Println(aocClient.Calendar())
	},
}

func init() {
	rootCmd.AddCommand(calendarCmd)
}
