package main

import (
	"fmt"

	"github.com/jskirchmeier/explore/adventure"

	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [path to file to validate",
	Short: "validate a explore yaml file",
	//Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(title)

		if len(args) < 1 {

			fmt.Println("Name of the adventure validate must be defined")
			return
		}
		fmt.Println("Validating adventure : ", args[0])

		a := adventure.New(args[0])

		if a == nil {

			fmt.Println("no adventure by the name " + args[0] + " was found")
			return
		}
		// TODO add validation code

	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
