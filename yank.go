package main

import (
	"bufio"
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
	fmt.Println("♥  Compiling hk-helpers...")

	inputFile := "src/_hk.css"
	getPurpleBuffer := readPurple()
	ss := css.Parse(getPurpleBuffer)
	rules := ss.GetCSSRuleList()

	file, err := os.Open(inputFile)
	checkErr(err)
	defer file.Close()

	outputFile, err := os.Create("src/_hk--compiled.css")
	checkErr(err)
	defer outputFile.Close()

	selectorCount := 0

	outputFile.WriteString("/* Don't edit this file directly */" + "\n")
	outputFile.WriteString("/* See the docs at https://github.com/heroku/purple3/blob/master/readme.md */" + "\n")

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
	color.Green("✓  Created /src/_hk--compiled.css")
}

func sanitize(input string) string {
	r := strings.NewReplacer(";", "", "  @yank ", "")
	result := r.Replace(input)
	return result
}

func readPurple() string {
	purpleFile := "css/purple3.css"

	dat, err := ioutil.ReadFile(purpleFile)
	checkErr(err)
	purpleFileBuffer := string(dat)
	return purpleFileBuffer
}
