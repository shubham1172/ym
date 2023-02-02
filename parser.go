package main

import "gopkg.in/yaml.v3"

// parseConfiguration takes a contents of a YAML file and returns a slice of YmEntry.
func parseConfiguration(data []byte) ([]YmEntry, error) {
	var entries []YmEntry
	err := yaml.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// parseFile takes a contents of a YAML file and returns a map of the contents.
func parseFile(data []byte) (map[string]interface{}, error) {
	var values map[string]interface{}
	err := yaml.Unmarshal(data, &values)
	if err != nil {
		return nil, err
	}
	return values, nil
}
