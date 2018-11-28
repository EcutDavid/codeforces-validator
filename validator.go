// TODO: add flags
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	defaultInSuffix     = ".in"
	defaultOutSuffix    = ".out"
	defaultCaseSplitter = "*****new case*****\n"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Solution path and test case prefix are needed as arguments")
	}
	filePath, casePrefix := os.Args[1], os.Args[2]
	inPath, outPath := casePrefix+defaultInSuffix, casePrefix+defaultOutSuffix

	in, err := ioutil.ReadFile(inPath)
	if err != nil {
		log.Fatal("something wrong with reading", inPath, err)
	}
	out, err := ioutil.ReadFile(outPath)
	if err != nil {
		log.Fatal("something wrong with reading", outPath, err)
	}
	cases := strings.Split(string(in), defaultCaseSplitter)
	results := strings.Split(strings.Replace(string(out), "\r\n", "\n", -1), defaultCaseSplitter)
	if len(cases) != len(results) {
		log.Fatalf("There are %d test cases, but %d expected results", len(cases), len(results))
	}

	// Validations.
	for i := 0; i < len(cases); i++ {
		fmt.Println("Running test case", i+1)
		fmt.Printf("The input is:\n%s\n", cases[i])

		cmd := exec.Command("go", "run", filePath)
		cmd.Stdin = strings.NewReader(cases[i])
		stdResult := bytes.Buffer{}
		cmd.Stderr = os.Stdout
		cmd.Stdout = &stdResult
		if err = cmd.Run(); err != nil {
			log.Fatal(err)
		}

		// Normalization
		if len(stdResult.Bytes()) > 0 {
			// New line in ASCII is 10
			if stdResult.Bytes()[stdResult.Len()-1] == 10 {
				if len(results[i]) > 0 && results[i][len(results[i])-1] != 10 {
					results[i] = results[i] + "\n"
				}
			}
			if stdResult.Bytes()[stdResult.Len()-1] != 10 {
				if len(results[i]) > 0 && results[i][len(results[i])-1] == 10 {
					results[i] = results[i][0 : len(results[i])-1]
				}
			}
		}

		if stdResult.String() != results[i] {
			fmt.Printf("Failed\nexpected result:\n%s\nbut get:\n%s", results[i], stdResult.String())
		} else {
			fmt.Printf("Passed\n\n")
		}
	}
}
