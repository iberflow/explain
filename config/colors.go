package config

const ArgColor = "#00ffae"
const FlagColor = "#fff000"

func ColorArg(arg string, flag bool) string {
	c := ArgColor

	if flag {
		c = FlagColor
	}

	arg = `[` + c + `]` + arg + `[white]`

	return arg
}