package main

import (
	"fmt"
	"os"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/commands"
)

// The version string is built into the binary with LDFLAGS
// -ldflags "-X main.version=$VERSION"
// To ensure clarity when building development versions the
// string below defaults to this value.
var version string = "v0.0.0-devel"

func main() {

	rootCommand := commands.CreateCommands(version)

	if err := rootCommand.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", commands.CliName, err.Error())
		_, _ = fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", commands.CliName)
		os.Exit(1)
	}
}
