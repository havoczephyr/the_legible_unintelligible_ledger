package main

import "regexp"

func BuildReplacers() []Replacer {
	var replacers []Replacer
	var regex *regexp.Regexp

	regex, _ = regexp.Compile(`(\[)UA(\])`)
	replacers = append(replacers, Replacer{regex, "Unintelligible:Articulation"})

	regex, _ = regexp.Compile(`(\[)UC(\])`)
	replacers = append(replacers, Replacer{regex, "Unintelligible:Crosstalk"})

	regex, _ = regexp.Compile(`(\[)BS(\])`)
	replacers = append(replacers, Replacer{regex, "BackgroundSpeech"})

	regex, _ = regexp.Compile(`(<[/]*)RN([^>]*>)`)
	replacers = append(replacers, Replacer{regex, "RedactedName"})

	regex, _ = regexp.Compile(`(<[/]*)RSE([^>]*>)`)
	replacers = append(replacers, Replacer{regex, "RedactedSensitive"})

	regex, _ = regexp.Compile(`(<[/]*)RSP([^>]*>)`)
	replacers = append(replacers, Replacer{regex, "RedactedSpeaker"})

	return replacers
}
