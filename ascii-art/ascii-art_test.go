/* Package contains the functions required to compare the command line output
with the expected result stored in a file.

It also makes two additional checks. (empty input and only "\n")

https://github.com/nik-don/01-founders/tree/main/ascii-art
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

// All additional test cases will be read from this file
var fileName string = "test-cases.txt"

type testData struct {
	arg      string
	expected string
}

var testCases = []testData{
	{"", ""},            // test case for empty argument ""
	{"\\n", "$" + "\n"}, // test case for only "\n"
}

// TestAsciiArt will test the output of the main.go program to match the desired representation in Ascii Art characters
func TestAsciiArt(t *testing.T) {
	// Log on/off
	logVisually := true

	// add additional testcases from the file into the testCases
	addTestCases(fileName)

	t.Parallel()
	for _, test := range testCases {

		cmd := "go run main.go " + escapedString(test.arg) + " | cat -e"

		output, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			t.Logf("try checking the argument in the test-case file")
			t.Fatal(err)
		}

		if logVisually {
			t.Logf("output for %s is \n%s", test.arg, output)
			t.Log()
			t.Logf("expected output for %s is \n%s", test.arg, test.expected)
		}

		if string(output) != (test.expected) {
			t.Errorf("For argument '%v' \nexpected \n%s\nbut got \n%s", test.arg, test.expected, string(output))
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

// escapedString is used to pass a string that is already escaped as an argument into the command line
// example:	"1a\"#F"
// becomes:	"1a"#F"
func escapedString(s string) string {
	s = fmt.Sprint(s)
	return ("\"" + s + "\"")
}
