// Package main provides the entry point for the transformation CLI tool.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cekrem/go-transform/internal/processor"
)

func main() {
	transformerName := flag.String("transformer", "passthrough", "name of the transformer to use")
	flag.Parse()

	// Get the executable's directory.
	execPath, err := os.Executable()
	if err != nil {
		log.Printf("Failed to get executable path: %v\n", err)
		os.Exit(1)
	}
	execDir := filepath.Dir(execPath)

	proc := processor.NewProcessor()

	// Load plugins from the plugins directory relative to the executable.
	pluginsDir := filepath.Join(execDir, "plugins")
	plugins, err := filepath.Glob(filepath.Join(pluginsDir, "*.so"))
	if err != nil || len(plugins) == 0 {
		log.Printf("Failed to list plugins: %v\n", err)
		os.Exit(1)
	}

	for _, plugin := range plugins {
		if err := proc.LoadPlugin(plugin); err != nil {
			log.Printf("Failed to load plugin %s: %v\n", plugin, err)

			continue
		}
	}

	// Read input from stdin, transform and print to stdout.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Bytes()

		output, err := proc.Process(*transformerName, input)
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		fmt.Println(string(output))
	}
}
