package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const RootDescription = "SICONV API Server"

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "server",
		Long: RootDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(2)
		},
	}

	rootCmd.AddCommand(commandStart())
	rootCmd.AddCommand(commandVersion())

	return rootCmd
}

func main() {
	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
