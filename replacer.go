package main

import "regexp"

type Replacer struct {
	Regex       *regexp.Regexp
	Replacement string
}
