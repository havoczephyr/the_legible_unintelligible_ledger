package main

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
)

type Replacer struct {
	Regex       regexp.Regexp
	Replacement string
}

const JSON_TAGS string = `{
	"[UA]" : "[Unintelligible:Articulation]",
	"[UC]" : "[Unintelligible:Crosstalk]",
	"[BS]" : "[BackgroundSpeech]",
	"RN" : "RedactedName",
	"RSE" : "RedactedSensitive",
	"RSP" : "RedactedSpeaker"
	}`

func readTags() map[string]string {
	var result map[string]string

	json.Unmarshal([]byte(JSON_TAGS), &result)
	return result
}

func main() {
	regexTags, _ := regexp.Compile(`(\[[A-Z]+\]|<([A-Z][^>]+)>|<\/([A-Z][^>]+)>)`)

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
		txtline := scanner.Text()
		regexTags.ReplaceAllString()
	}
	tagMap := readTags()

}
