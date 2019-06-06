package main

import (
	"fmt"
	"os"

	"github.com/jskirchmeier/explore/adventure"
	"github.com/jskirchmeier/explore/cli"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load [path to file]",
	Short: "load and start (continue) an Exploration",
	//Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(title)

		if len(args) < 1 {
			fmt.Println("must pass in the adventure you want to do")
			os.Exit(1)
		}
		a := adventure.New(args[0])
		if a == nil {
			fmt.Println("No adventure with the name " + args[0] + " is available")
			os.Exit(2)
		}
		runner := cli.Runner{}
		runner.Run(a) // blocks

	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
