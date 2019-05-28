package main

import (
	"fmt"
	"github.com/jskirchmeier/explore"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [path to file to validate",
	Short: "validate a explore yaml file",
	//Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(title)

		if len(args) < 1 {

			fmt.Println("Path to file to validate must be defined")
			return
		}
		fmt.Println("Validating adventure : ", args[0])

		a, err := explore.NewAdventure(args[0])

		if err != nil {
			fmt.Println(err)
			return
		}
		a.Validate()

	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
