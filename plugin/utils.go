package plugin

import (
	"io/ioutil"
	"os"
)

func readStringOrFile(input string) (string, bool, error) {
	if len(input) > 255 {
		return input, false, nil
	}
	// Check if input is a file path
	if _, err := os.Stat(input); err != nil && os.IsNotExist(err) {
		// No file found => use input as result
		return input, false, nil
	} else if err != nil {
		return "", false, err
	}
	result, err := ioutil.ReadFile(input)
	if err != nil {
		return "", true, err
	}
	return string(result), true, nil
}
