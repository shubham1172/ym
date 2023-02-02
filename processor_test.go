package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const yamlWithArrayRoot = `
- name: foo
  versions:
	- firstValue
	- secondValue
  metadata:
	ttl: 30
- name: bar
  versions:
  	- thirdValue
	- fourthValue
  metadata:
	ttl: 60`

func TestLookup(t *testing.T) {
	tcs := []struct {
		inputStr string
		path     string
		value    interface{}
		err      error
	}{
		{
			inputStr: yamlWithArrayRoot,
			path:     "[0].versions[0]",
			value:    "firstValue",
			err:      nil,
		},
		{
			inputStr: yamlWithArrayRoot,
			path:     "[0].metadata.ttl",
			value:    30,
			err:      nil,
		},
		{
			inputStr: yamlWithArrayRoot,
			path:     "[1].name",
			value:    "bar",
			err:      nil,
		},
		{
			inputStr: yamlWithArrayRoot,
			path:     "[1].versions[1]",
			value:    "fourthValue",
			err:      nil,
		},
		{
			inputStr: yamlWithArrayRoot,
			path:     "[1]",
			value: map[string]interface{}{
				"name": "bar",
				"versions": []interface{}{
					"thirdValue",
					"fourthValue",
				},
				"metadata": map[string]interface{}{
					"ttl": 60,
				},
			},
			err: nil,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.path, func(t *testing.T) {
			var input interface{}
			err := yaml.Unmarshal([]byte(tc.inputStr), &input)
			assert.NoError(t, err)

			value, err := lookupNode(input, tc.path)
			if err != tc.err {
				t.Errorf("got err %v, want %v", err, tc.err)
			}
			if value != tc.value {
				t.Errorf("got value %v, want %v", value, tc.value)
			}
		})
	}
}
