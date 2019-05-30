package main

import (
	"fmt"
	"github.com/jskirchmeier/explore"
	"github.com/jskirchmeier/explore/cli"
	"github.com/spf13/cobra"
	"os"
)

var loadCmd = &cobra.Command{
	Use:   "load [path to file]",
	Short: "load and start (continue) an Exploration",
	//Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(title)

		// the root run is the CLI runner
		// to run the server use that command
		if len(args) < 1 {
			fmt.Println("must pass in the adventure you want to do")
			os.Exit(1)
		}
		a, err := explore.NewAdventure(args[0])
		if err != nil {
			fmt.Println("Unable to start", err.Error())
			os.Exit(2)
		}

		runner := cli.Runner{}
		runner.Run(a) // blocks

	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
