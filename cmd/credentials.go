package cmd

import (
	"github.com/spf13/cobra"
	"github.com/teodorpopa/advent-of-code-go/internal"
	"log"
)

var cookieStr string

var credentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Create or update the credentials file",
	Long:  `Create or update the credentials file used to communicate with the AoC website.`,
	Run: func(cmd *cobra.Command, args []string) {

		credentials := internal.CredentialsManager{
			Cookie: cookieStr,
		}

		err := credentials.Update()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(credentialsCmd)

	credentialsCmd.Flags().StringVarP(&cookieStr, "cookie", "c", "", "Cookie string")
	err := credentialsCmd.MarkFlagRequired("cookie")

	if err != nil {
		log.Fatal(err)
		return
	}
}
