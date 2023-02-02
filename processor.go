package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mikefarah/yq/v4/pkg/yqlib"
)

const (
	opUpdate operation = "update"
	opDelete operation = "delete"
)

var (
	evaluator = yqlib.NewStringEvaluator()
	encoder   = yqlib.NewYamlEncoder(2, false, yqlib.YamlPreferences{})
	decoder   = yqlib.NewYamlDecoder(yqlib.YamlPreferences{})
)

type operation string

// YmOperation represents a single operation to be applied to an input file.
type YmOperation struct {
	Operation operation   `yaml:"op"`
	Path      string      `yaml:"path"`
	Value     interface{} `yaml:"value,omitempty"`
}

// YmEntry represents an input file, the operations to be applied to it,
// and the output file to write the results to.
type YmEntry struct {
	Input      string        `yaml:"input"`
	Output     string        `yaml:"output,omitempty"`
	Operations []YmOperation `yaml:"operations"`
}

// process takes a slice of YmEntry and processes them individually.
// In case of an error, it returns immediately.
func process(opsFilePath string, entries []YmEntry) error {
	for _, entry := range entries {
		if err := processEntry(opsFilePath, entry); err != nil {
			return err
		}
	}
	return nil
}

// processEntry takes a single YmEntry, applies the operations to the file,
// writes the results to the output file, and returns an error if any.
func processEntry(opsFilePath string, entry YmEntry) error {
	data, err := os.ReadFile(getPath(opsFilePath, entry.Input))
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", entry.Input, err)
	}

	// Create a yq expression from the operations.
	expr := ""

	for _, op := range entry.Operations {
		switch op.Operation {
		case opUpdate:
			jsonBytes, err := json.Marshal(op.Value)
			if err != nil {
				return err
			}
			expr += "|" + fmt.Sprintf("%s=%s", op.Path, jsonBytes)
		case opDelete:
			expr += "|" + fmt.Sprintf("del(%s)", op.Path)
		default:
			return fmt.Errorf("unknown operation %s", op.Operation)
		}
	}

	expr = expr[1:] // Remove the first "|".

	log.Printf("[INFO] Applying expression %s\n", expr)

	// Apply the expression to the file.
	s, err := evaluator.Evaluate(expr, string(data), encoder, decoder)
	if err != nil {
		return err
	}

	if entry.Output == "" {
		log.Printf("[WARN] No output file specified, input file will be overwritten")
		entry.Output = entry.Input
	}
	return os.WriteFile(getPath(opsFilePath, entry.Output), []byte(s), 0644)
}

// getPath returns the filePath as an absolute path.
// If filePath is relative, it's resolved relative to the opsFilePath.
func getPath(opsFilePath string, filePath string) string {
	if filepath.IsAbs(filePath) {
		return filePath
	}
	absPath, err := filepath.Abs(filepath.Dir(opsFilePath))
	if err != nil {
		return filepath.Join(filepath.Dir(opsFilePath), filePath)
	} else {
		return filepath.Join(absPath, filePath)
	}
}
