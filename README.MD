# The Legible Unintelligible Ledger!
### by Giovanni D'Amico

*"The slightly intelligible solution!"*

`t_lul` is another small .exe I use because I'm lazy and don't like typing multiple `[Unintelligible:Articulation]` in a single sentence.

Utlizes regular expressions to sift through lines of `.tsv` files and replaces shorthand with the appropriate forms.

## How to use

stick `t_lul` in your path and call it in a folder with a `pre-curated.tsv` in it. it'll generate a `curated.tsv` with everything expanded appropriately. thats it! 

**`t_lul` will panic if `pre-curated.tsv` is not present in the working directory.**

## Expansion Key
Shorthand | Expanded
--|--
`[UA]` | `[Unintelligible:Articulation]`
`[UC]` | `[Unintelligible:Crosstalk]`
`[BS]` | `[BackgroundSpeech]`
`<RSE>`, `</RSE>` | `<RedactedSensitive>`,`</RedactedSensitive>`
`<RSE time="HH:MM:SS.sss-HH:MM:SS.sss">` | `<RedactedSensitive time="HH:MM:SS.sss-HH:MM:SS.sss">`
`<RN>`, `</RN>` | `<RedactedName>`, `</RedactedName>`
`<RN time="HH:MM:SS.sss-HH:MM:SS.sss">` |`<RedactedName time="HH:MM:SS.sss-HH:MM:SS.sss">`
`<RSP>`, `</RSP>` | `<RedactedSpeaker>`, `</RedactedSpeaker>` 