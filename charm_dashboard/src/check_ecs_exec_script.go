package main

import (
	"embed"
	"fmt"
	"os"
	"path"
)

//go:embed check-ecs-exec.sh
var script embed.FS

func installCheck() error {

	// Read the embedded script
	scriptDir := "/tmp"
	scriptName := "check-ecs-exec.sh"

	fullPath := path.Join(scriptDir, scriptName)

	// If the file exists then we don't need to install it
	if _, err := os.Stat(fullPath); err == nil {
		return nil
	}

	data, err := script.ReadFile(scriptName)
	if err != nil {
		return fmt.Errorf("Error reading embedded script %s: %w", scriptName, err)
	}

	// Write the script to a file
	err = os.WriteFile(fullPath, data, 0755)

	return err
}
