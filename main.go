package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("The Legible Unintelligible Ledger - Giovanni D'Amico \n Expanding pre-curated.tsv into curated.tsv...")
	replacers := BuildReplacers()

	file, err := os.Open("pre-curated.tsv")
	if err != nil {
		fmt.Printf("Could not open or find a pre-curated.tsv file. -%v", err)
	}
	defer file.Close()

	curated, err := os.Create("curated.tsv")
	if err != nil {
		fmt.Printf("Could not create a curated.tsv file. -%v", err)
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
	fmt.Println("\n🔥 Process Complete. Have a nice day 🔥")
}
