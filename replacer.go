package main

import "regexp"

//struct used by BuildReplacers().
type Replacer struct {
	Regex       *regexp.Regexp
	Replacement string
}
