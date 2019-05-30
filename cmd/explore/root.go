package main

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

}

var rootCmd = &cobra.Command{
	Use:   "explore",
	Short: title,
	Long:  title + directions,
}
