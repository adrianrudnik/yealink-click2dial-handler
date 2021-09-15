package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	rootCmd.SetOut(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setup() {
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(dialCmd)
	rootCmd.AddCommand(resetCmd)
}

func init() {
	setup()
}

var rootCmd = &cobra.Command{
	Use:   "yealink-url-scheme-handler",
	Short: "Connects a single desktop phone with the tel and callto URL scheme handlers of the OS",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
