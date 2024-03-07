package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func printVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("%s version %s\n", cliName, ver)
	return nil
}
