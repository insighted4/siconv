package main

import (
	"fmt"
	"runtime"

	"github.com/insighted4/siconv/version"
	"github.com/spf13/cobra"
)

func commandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version and exit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RootDescription)
			fmt.Printf("Version: %s\n", version.Version)
			fmt.Printf("Go Version: %s\n", runtime.Version())
			fmt.Printf("Go OS/ARCH: %s %s\n", runtime.GOOS, runtime.GOARCH)
		},
	}
}
