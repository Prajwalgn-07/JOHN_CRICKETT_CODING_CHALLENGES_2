package handlers

import (
	"bufio"
	"os"
)

func IsHostForbidden(host string) (bool, error) {
	// open the file using Open() function from os library
	file, err := os.Open("forbidden-hosts.txt")
	if err != nil {
		return false, err
	}
	defer file.Close()

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == host {
			return true, nil
		}
	}
	// check for the error that occurred during the scanning
	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}
