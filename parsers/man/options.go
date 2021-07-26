package man

import (
	"regexp"
	"strings"
)

var paramPattern = regexp.MustCompile(`(?i)<(.*)>`)

type Option struct {
	Name          string
	NameIndicator string

	ShortName          string
	ShortNameIndicator string

	Parameter   string
	Description string
}

func (o *Option) String() string {
	name := o.Name
	if name == o.ShortName {
		name = o.NameIndicator + name
	} else {
		name = o.ShortNameIndicator + o.ShortName + ", " + o.NameIndicator + name
	}

	return name
}

func NewOption(name, description string) *Option {
	return buildOption(name, description)
}

func buildOption(name string, description string) *Option {
	var args []string
	var param string

	// extract <parameter> and remove it from name
	if paramPattern.MatchString(name) {
		param = strings.Trim(paramPattern.FindString(name), "<>")
		name = paramPattern.ReplaceAllString(name, "")
	}

	// split into multiple args and trim
	for _, arg := range strings.Split(name, ",") {
		args = append(args, strings.TrimSpace(arg))
	}

	opt := &Option{}
	opt.Parameter = param
	opt.Description = description

	// handle arg having a short and a long name
	if len(args) == 2 {
		if len(args[0]) < len(args[1]) {
			opt.Name = args[1]
			opt.ShortName = args[0]
		}

		if len(args[1]) < len(args[0]) {
			opt.Name = args[0]
			opt.ShortName = args[1]
		}
	}

	// if only a single arg defined, use same name for both the full and the short names
	if len(args) == 1 {
		opt.Name = args[0]
		opt.ShortName = args[0]
	}

	opt.NameIndicator = getIndicator(opt.Name)
	opt.ShortNameIndicator = getIndicator(opt.ShortName)

	opt.Name = strings.Trim(opt.Name, "-")
	opt.ShortName = strings.Trim(opt.ShortName, "-")

	return opt
}

func getIndicator(name string) string {
	if strings.Contains(name, "--") {
		return "--"
	}

	if strings.Contains(name, "-") {
		return "-"
	}

	return ""
}
