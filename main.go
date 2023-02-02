package main

import (
	"flag"
	"log"
	"os"
)

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

	if err := process(entries); err != nil {
		log.Fatalf("error processing: %v", err)
	}
}
