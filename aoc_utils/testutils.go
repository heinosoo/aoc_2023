package aoc_utils

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path"
	"testing"
	"time"
)

func TestPart(t *testing.T, testee func(<-chan string, chan string), pathSuffix string) {
	log.Println("\n\nTesting: ", pathSuffix)
	outputPath := "output_" + pathSuffix
	expectedPath := "expected_" + pathSuffix
	for _, testFileName := range getInputFileNames() {
		t.Run(testFileName, func(t *testing.T) {
			input := readFile("input/" + testFileName)
			output := make(chan string, 10)
			withTiming(testFileName, func() {
				testee(input, output)
			})
			writeFile(path.Join(outputPath, testFileName), output)

			// Skip if there is no expected output
			if _, err := os.Stat(path.Join(expectedPath, testFileName)); os.IsNotExist(err) {
				log.Println(testFileName, "no expected output")
				t.Skip()
			}

			compareFiles(path.Join(outputPath, testFileName), path.Join(expectedPath, testFileName))
		})
	}
}

func readFile(filename string) <-chan string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := make(chan string)
	scanner := bufio.NewScanner(file)
	go func() {
		defer file.Close()
		defer close(lines)

		for scanner.Scan() {
			lines <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	return lines
}

func writeFile(filename string, output chan string) {
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	for line := range output {
		_, e := file.WriteString(line + "\n")
		if e != nil {
			panic(e)
		}
	}
}

func withTiming(prefix string, doSomething func()) {
	t := time.Now()
	doSomething()
	log.Println(prefix, time.Since(t))
}

func compareFiles(s1, s2 string) {
	// call diff
	cmd := exec.Command("diff", s1, s2)
	// wait for it to finish
	diff, _ := cmd.Output()
	if len(diff) > 0 {
		log.Fatal("Files differ:\n", string(diff))
	}
}

func getInputFileNames() []string {
	var filenames []string

	// Default to all files in the "input" directory
	files, err := os.ReadDir("input")
	if err != nil {
		log.Fatal(err)
	}

	// Append filenames to the slice
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}

	log.Println("Found", len(filenames), "input file(s)")
	return filenames
}
