package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/mikefarah/yq/v4/pkg/yqlib"
	"gopkg.in/op/go-logging.v1"
	"gopkg.in/yaml.v3"
)

// parseConfiguration takes a contents of a YAML file and returns a slice of YmEntry.
func parseConfiguration(data []byte) ([]YmEntry, error) {
	var entries []YmEntry
	err := yaml.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// getFileContent returns the contents of a file.
// If the file is not provided, it will try to read from stdin.
func getFileContent(f string) ([]byte, error) {
	if f == "" {
		fi, err := os.Stdin.Stat()
		// Check if stdin is a pipe or a terminal
		if err != nil || (fi.Mode()&os.ModeCharDevice) != 0 {
			flag.Usage()
			os.Exit(1)
		}
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(f)
}

func main() {
	var file string
	flag.StringVar(&file, "file", "", "path to the operations file (or specify via stdin)")
	flag.Parse()

	data, err := getFileContent(file)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	entries, err := parseConfiguration(data)
	if err != nil {
		log.Fatalf("error parsing YAML: %v", err)
	}

	// Discard the yqlib logger!
	discardBackend := logging.AddModuleLevel(logging.NewLogBackend(io.Discard, "", 0))
	yqlib.GetLogger().SetBackend(discardBackend)

	if err := process(file, entries); err != nil {
		log.Fatalf("error processing: %v", err)
	}
}
