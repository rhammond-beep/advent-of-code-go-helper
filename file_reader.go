package util

import (
	"bufio"
	"os"
	"strings"
)

/*
Reads file line by line, converting each as a string and appending
to return slice
*/
func ReadChallengeInput(filepath string) (fileContents []string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContents = append(fileContents, scanner.Text())
	}
	return
}

/*
Reads entire UTF-8 encoded file as a one contiguous string
*/
func ReadChallengeInputContigious(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	return sb.String()
}
