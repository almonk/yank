package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"sort"

	"github.com/almonk/css"
	"github.com/fatih/color"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func main() {
	// Params
	inputFilePtr := flag.String("input", "myclasses.css", "File where your @yank rules are defined")
	definitionsFilePtr := flag.String("definitions", "mycss.css", "File that @yank uses to expand your classes")
	outputFilePtr := flag.String("output", "myclasses--compiled.css", "CSS file with compiled @yank rules")

	flag.Parse()

	fmt.Println("♥  Compiling @yank...")

	inputFile := *inputFilePtr
	getPurpleBuffer := readPurple(*definitionsFilePtr)
	ss := css.Parse(getPurpleBuffer)
	rules := ss.GetCSSRuleList()

	file, err := os.Open(inputFile)
	checkErr(err)
	defer file.Close()

	outputFile, err := os.Create(*outputFilePtr)
	checkErr(err)
	defer outputFile.Close()

	selectorCount := 0

	outputFile.WriteString("/* This file was compiled with @yank */" + "\n")
	outputFile.WriteString("/* See the docs at https://github.com/almonk/yank */" + "\n")

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, ".") {
			sanitizedLine := sanitize(line)
			outputFile.WriteString(sanitizedLine + "\n")
			selectorCount++
			outputFile.Sync()
		}

		if strings.HasPrefix(line, "  ") {
			if strings.HasPrefix(line, "  @yank ") {
				// Lets expand what we're yanking;
				sanitizedLine := sanitize(line)

				for _, rule := range rules {
					if rule.Style.SelectorText == sanitizedLine {
						// Found the matching rules
						allRules := []string{}

						for _, v := range rule.Style.Styles {
							writeRule := fmt.Sprintf("%s:%s;", v.Property, v.Value)
							allRules = append(allRules, writeRule)
						}

						sort.Strings(allRules)

						for _, w := range allRules {
							outputFile.WriteString(w + "\n")
							outputFile.Sync()
						}

					}
				}
			} else {
				// Just print the line out
				outputFile.WriteString(line + "\n")
				outputFile.Sync()
			}
		}

		if strings.HasPrefix(line, "}") {
			outputFile.WriteString("}\n")
			outputFile.Sync()
		}

	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	resultSelectorCount := fmt.Sprintf("✓  Found %v classes", selectorCount)
	color.Green(resultSelectorCount)

	successMsg := fmt.Sprintf("✓  Created %s", *outputFilePtr)
	color.Green(successMsg)
}

func sanitize(input string) string {
	r := strings.NewReplacer(";", "", "  @yank ", "")
	result := r.Replace(input)
	return result
}

func readPurple(inputFile string) string {
	purpleFile := inputFile

	dat, err := ioutil.ReadFile(purpleFile)
	checkErr(err)
	purpleFileBuffer := string(dat)
	return purpleFileBuffer
}
