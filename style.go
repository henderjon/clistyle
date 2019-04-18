package clistyle

import (
	"strings"
)

// These are the bits to use in setting the flag
const (
	esc       = "\u001b[" // "\033", byte(27)
	reset     = ";0m"
	bold      = ";1"
	underline = ";4"
	reverse   = ";7"
	black     = ";30"
	red       = ";31"
	green     = ";32"
	yellow    = ";33"
	blue      = ";34"
	magenta   = ";35"
	cyan      = ";36"
	white     = ";37"
	blackBG   = ";40"
	redBG     = ";41"
	greenBG   = ";42"
	yellowBG  = ";43"
	blueBG    = ";44"
	magentaBG = ";45"
	cyanBG    = ";46"
	whiteBG   = ";47"
	None      = 0
	Bold      = 1 << iota
	Underline
	Reverse
	Bright
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BlackBG
	RedBG
	GreenBG
	YellowBG
	BlueBG
	MagentaBG
	CyanBG
	WhiteBG
)

var texts = map[int]string{
	Reverse:   reverse,
	Underline: underline,
	Bold:      bold,
}
var foregrounds = map[int]string{
	Black:   black,
	Red:     red,
	Green:   green,
	Yellow:  yellow,
	Blue:    blue,
	Magenta: magenta,
	Cyan:    cyan,
	White:   white,
}
var backgrounds = map[int]string{
	BlackBG:   blackBG,
	RedBG:     redBG,
	GreenBG:   greenBG,
	YellowBG:  yellowBG,
	BlueBG:    blueBG,
	MagentaBG: magentaBG,
	CyanBG:    cyanBG,
	WhiteBG:   whiteBG,
}

// Style decorates value according to the options passed via flag. It only allows three settings
// [TEXT;FOREGROUND;BACKGROUNDm. Bits are or'ed together to control the style.
func Style(value string, flag int) string {
	var (
		s strings.Builder
		tx,
		fg,
		bg string
	)
	for b, s := range texts {
		if has(flag, b) {
			tx = s
		}
	}
	for b, s := range foregrounds {
		if has(flag, b) {
			fg = s
		}
	}
	for b, s := range backgrounds {
		if has(flag, b) {
			bg = s
		}
	}

	if len(tx)+len(fg)+len(bg) > 0 {
		s.WriteString(esc)
		s.WriteString(tx)
		s.WriteString(fg)
		s.WriteString(bg)
		s.WriteString("m") // ? ... so strange
		s.WriteString(value)
		s.WriteString(esc)
		s.WriteString(reset)
	} else {
		return value
	}

	return s.String()
}

// https://yourbasic.org/golang/bitmask-flag-set-clear/
// func set(b, flag int) int    { return b | flag }
// func clear(b, flag int) int  { return b &^ flag }
// func toggle(b, flag int) int { return b ^ flag }
func has(b, flag int) bool { return b&flag != 0 }
