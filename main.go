package main

import (
	"bufio"
	"os"
)

func main() {

	replacers := BuildReplacers()

	file, err := os.Open("pre-curated.tsv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	curated, err := os.Create("curated.tsv")
	if err != nil {
		panic(err)
	}
	defer curated.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		textline := scanner.Text()
		for _, replacer := range replacers {
			if replacer.Regex.MatchString(textline) {
				replacement := "${1}" + replacer.Replacement + "${2}"
				textline = replacer.Regex.ReplaceAllString(textline, replacement)
			}

		}
		curated.WriteString(textline + "\n")
	}
}
