/* Package contains the functions required to compare the command line output
with the expected result stored in a file.

https://github.com/nik-don/01-founders/tree/main/quad
*/
package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"piscine"
	"strconv"
	"strings"
	"testing"
)

// All additional test cases will be read from this file
var fileName string = "test-cases.txt"

type testData struct {
	q        string
	x        int
	y        int
	expected string
}

var testCases = []testData{}

// TestQuad will test the output of the main.go program to match the desired quadrangle result
func TestQuad(t *testing.T) {
	// add additional testcases from the file into the testCases
	addTestCases(fileName)

	for _, test := range testCases {
		out := runTest(test.q, test.x, test.y, test.expected)

		if out != test.expected {
			t.Errorf("For argument Quad%v (%v,%v) \n\033[1;34mexpected \n%s\033[0m\n\033[1;31mbut got \n%s\033[0m", test.q, test.x, test.y, test.expected, string(out))
			panic("err")
		} else {
			t.Logf("output for Quad%s (%v,%v) is \n\033[1;32m%s\033[0m", test.q, test.x, test.y, out)
			t.Log("\n")
			t.Logf("expected output for Quad%s (%v,%v) is \n\033[1;32m%s\033[0m\n\033[1;36m------------------------------\033[0m\n", test.q, test.x, test.y, test.expected)

		}
	}
}

// runTest runs individual tests of the required function of the package piscine and returns its output
func runTest(q string, x, y int, exp string) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	switch q {
	case "A":
		piscine.QuadA(x, y)
	case "B":
		piscine.QuadB(x, y)
	case "C":
		piscine.QuadC(x, y)
	case "D":
		piscine.QuadD(x, y)
	case "E":
		piscine.QuadE(x, y)
	}

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}

// addTestCases will read the contents of the specified file, extract and add the desired output into the testCases
func addTestCases(fileName string) {
	args := []string{}
	tempStr := ""

	lines, err := readLines(fileName)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		if line[0] == '#' {
			if tempStr != "" {
				x, _ := strconv.Atoi(args[1])
				y, _ := strconv.Atoi(args[2])
				testCases = append(testCases, testData{args[0], x, y, tempStr})
				tempStr = ""
			}
			args = strings.Split(string(line[1:len(line)-1]), ",")

		} else if line[0] == '@' {
			x, _ := strconv.Atoi(args[1])
			y, _ := strconv.Atoi(args[2])
			testCases = append(testCases, testData{args[0], x, y, ""})

		} else {
			tempStr += line + "\n"
		}
		// Last test case(No more testcases in the file)
		if i == len(lines)-1 {
			x, _ := strconv.Atoi(args[1])
			y, _ := strconv.Atoi(args[2])
			testCases = append(testCases, testData{args[0], x, y, tempStr})
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
