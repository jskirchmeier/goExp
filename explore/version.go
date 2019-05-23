package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of explore!",
	Long:  "Print the version number of explore!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(title)
		fmt.Println("Version : ", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
