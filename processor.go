package main

import (
	"fmt"
	"os"
)

const (
	opReplace operation = "replace"
	opAdd     operation = "add"
	opRemove  operation = "remove"
)

type operation string

// YmOperation represents a single operation to be applied to a file.
type YmOperation struct {
	Operation operation   `yaml:"op"`
	Path      string      `yaml:"path"`
	Value     interface{} `yaml:"value,omitempty"`
}

// YmEntry represents a single file and the operations to be applied to it.
type YmEntry struct {
	File       string        `yaml:"file"`
	Operations []YmOperation `yaml:"operations"`
}

// process takes a slice of YmEntry and processes them individually.
// In case of an error, it returns immediately.
func process(entries []YmEntry) error {
	for _, entry := range entries {
		if err := processEntry(entry); err != nil {
			return err
		}
	}
	return nil
}

// processEntry takes a single YmEntry and applies the operations to the file.
// It overwrites the file with the new contents and returns an error if any.
func processEntry(entry YmEntry) error {
	data, err := os.ReadFile(entry.File)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", entry.File, err)
	}

	val, err := parseFile(data)
	if err != nil {
		return fmt.Errorf("error parsing file %s: %v", entry.File, err)
	}

	for _, op := range entry.Operations {
		switch op.Operation {
		case opReplace:
			if err := replaceNode(val, op.Path, op.Value); err != nil {
				return fmt.Errorf("error replacing value for %s: %v", op.Path, err)
			}
		case opAdd:
			if err := addNode(val, op.Path, op.Value); err != nil {
				return fmt.Errorf("error adding value for %s: %v", op.Path, err)
			}
		case opRemove:
			if err := removeNode(val, op.Path); err != nil {
				return fmt.Errorf("error removing value for %s: %v", op.Path, err)
			}
		default:
			return fmt.Errorf("unknown operation %s", op.Operation)
		}
	}

	return nil
}

// replaceNode takes a map of values, a path and a value and replaces the value at the path.
// If the path does not exist, it returns an error.
func replaceNode(val interface{}, path string, value interface{}) error {
	return nil
}

// addNode takes a map of values, a path and a value and adds the value at the path.
// If the path does not exist, it creates the path and adds the value.
// If the path already exists, it returns an error.
func addNode(val interface{}, path string, value interface{}) error {
	return nil
}

// removeNode takes a map of values and a path and removes the value at the path.
func removeNode(val interface{}, path string) error {
	return replaceNode(val, path, nil)
}

// lookupNode recursively traverses the map of values and returns the value at the path.
// The path is of the form foo.bar[0].baz.
func lookupNode(val interface{}, path string) (interface{}, error) {
	return nil, nil
}

// func lookupNode(val interface{}, path string) (interface{}, error) {
// 	ptr := val
// 	i := strings.Index(path, ".")
// 	if i == -1 {
// 		// no more segments, return the value
// 		return ptr[path], nil
// 	}
// 	// get the next segment
// 	segment := path[:i]
// 	if strings.Contains(segment, "[") {
// 		// handle array type
// 		i1 := strings.Index(segment, "[")
// 		if i1 == -1 {
// 			return nil, fmt.Errorf("invalid path %s, expected [", segment)
// 		}
// 		i2 := strings.Index(segment, "]")
// 		if i2 == -1 || i2 <= i1 {
// 			return nil, fmt.Errorf("invalid path %s, expected ] after [", segment)
// 		}
// 		if segment[i2+1:] != "" {
// 			return nil, fmt.Errorf("invalid path %s, nothing after ]", segment)
// 		}
// 		// get the array index
// 		idxStr := segment[i1+1 : i2]
// 		idx, err := strconv.Atoi(idxStr)
// 		if err != nil {
// 			return nil, fmt.Errorf("invalid path %s, expected integer after [", segment)
// 		}
// 		// get the array
// 		arr, ok := ptr[segment[:i1]].([]interface{})
// 		if !ok {
// 			return nil, fmt.Errorf("invalid path %s, expected array", segment)
// 		}
// 		// get the next value
// 		ptr = arr[idx]

// 	} else {
// 		// handle map type
// 	}
// }
