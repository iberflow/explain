package text

import (
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildCommandWithRegionsToString(t *testing.T) {
	cmdStr := "curl -sSL -a --basic 'https://install.larashed.com/linux'"
	command := args.NewCommand(cmdStr)

	basicOpt := &man.Option{
		Name:           "basic",
		NameIndicator:  "--",
		Alias:          "basic",
		AliasIndicator: "--",
		Parameter:      "",
		Description:    "",
	}
	appendOpt := &man.Option{
		Name:           "append",
		NameIndicator:  "--",
		Alias:          "a",
		AliasIndicator: "-",
		Parameter:      "",
		Description:    "",
	}
	silentOpt := &man.Option{
		Name:           "silent",
		NameIndicator:  "--",
		Alias:          "s",
		AliasIndicator: "-",
		Parameter:      "",
		Description:    "",
	}
	showErrorOpt := &man.Option{
		Name:           "show-error",
		NameIndicator:  "--",
		Alias:          "S",
		AliasIndicator: "-",
		Parameter:      "",
		Description:    "",
	}
	locOpt := &man.Option{
		Name:           "location",
		NameIndicator:  "--",
		Alias:          "L",
		AliasIndicator: "-",
		Parameter:      "",
		Description:    "",
	}

	list := &man.List{}
	list.Add(basicOpt)
	list.Add(appendOpt)
	list.Add(silentOpt)
	list.Add(showErrorOpt)
	list.Add(locOpt)

	expected := `curl -["0"][#00ffae]s[white][""]["0"][#00ffae]S[white][""]["0"][#00ffae]L[white][""] ["1"][#fff000]-a[white][""] ["2"][#fff000]--basic[white][""] [#00ffae]'https://install.larashed.com/linux'[white]`
	assert.Equal(t, expected, RenderCommand(command, list))
}
