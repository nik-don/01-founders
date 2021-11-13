/* Package contains the functions required to compare the command line output
with the expected result stored in a file.

https://github.com/nik-don/01-founders/tree/main/sudoku
*/
package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"testing"
)

// All additional test cases will be read from this file
var fileName string = "test-cases.txt"

const (
	InfoColor  = "\033[1;34m%s\033[0m"
	ErrorColor = "\033[1;31m%s\033[0m"
)

type testData struct {
	arg      string
	expected string
}

var testCases = []testData{}

// TestSudoku will test the output of the main.go program to match the solved sudoku output 
func TestSudoku(t *testing.T) {
	// add additional testcases from the file into the testCases
	addTestCases(fileName)

	for _, test := range testCases {

		cmd := "go run main.go " + (test.arg)

		output, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			t.Fatal(">>> TRY CHECKING the input argument in the", fileName, " file. For >>>> ", test.arg, " \n", err)
		}

		if string(output) != (test.expected) {
			t.Errorf("For argument '%v' \n\033[1;34mexpected \n%s\033[0m\n\033[1;31mbut got \n%s\033[0m", test.arg, test.expected, string(output))
			panic("err")
		} else {
			t.Logf("output for %s is \n\033[1;32m%s\033[0m", test.arg, output)
			t.Log("\n")
			t.Logf("expected output for %s is \n\033[1;32m%s\033[0m\n\033[1;36m------------------------------\033[0m\n", test.arg, test.expected)

		}

	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// addTestCases will read the contents of the specified file, extract and add the desired output into the testCases
func addTestCases(fileName string) {
	arg, tempStr := "", ""

	lines, err := readLines(fileName)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		if line[0] == '#' {
			if tempStr != "" {
				testCases = append(testCases, testData{arg, tempStr})
				tempStr = ""
			}
			arg = string(line[1 : len(line)-1])

		} else {
			tempStr += line + "\n"
		}
		// Last test case(No more testcases in the file)
		if i == len(lines)-1 {
			testCases = append(testCases, testData{arg, tempStr})
		}
	}
}
