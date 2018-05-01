package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const RootDescription = "SICONV Database Updater"

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "updater",
		Long: RootDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(2)
		},
	}

	rootCmd.AddCommand(commandSync())
	rootCmd.AddCommand(commandVersion())

	return rootCmd
}

func main() {
	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
