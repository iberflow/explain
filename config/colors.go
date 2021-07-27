package config

import "strings"

const ArgColor = "#00ffae"
const FlagColor = "#fff000"

func ColorArg(arg string) string {
	color := ArgColor

	if strings.Contains(arg, "-") {
		color = FlagColor
	}

	arg = `[` + color + `]` + arg + `[white]`

	return arg
}