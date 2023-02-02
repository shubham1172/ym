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

func main() {
	var file string
	flag.StringVar(&file, "file", "", "path to the operations file")
	flag.Parse()

	if file == "" {
		flag.Usage()
		os.Exit(1)
	}

	data, err := os.ReadFile(file)
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
