package man

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOption(t *testing.T) {
	r := buildOption("", "-m, --max-time <seconds>", "some description")

	assert.Equal(t, "max-time", r.Name)
	assert.Equal(t, "--", r.NameIndicator)
	assert.Equal(t, "m", r.Alias)
	assert.Equal(t, "-", r.AliasIndicator)
	assert.Equal(t, "seconds", r.Parameters[0])
	assert.Equal(t, "some description", r.Description)
}

func TestParseOption2(t *testing.T) {
	r := buildOption("", "\\fB-t \\fItimelimit\\fR\\fR", "some description")

	assert.Equal(t, "t", r.Name)
	assert.Equal(t, "-", r.NameIndicator)
	assert.Equal(t, "t", r.Alias)
	assert.Equal(t, "-", r.AliasIndicator)
	assert.Equal(t, "timelimit", r.Parameters[0])
	assert.Equal(t, "some description", r.Description)
}

func TestParseOption3(t *testing.T) {
	r := buildOption("", "\\fB-z \\fI<td>-attributes\\fR\\fR", "some description")

	assert.Equal(t, "z", r.Name)
	assert.Equal(t, "-", r.NameIndicator)
	assert.Equal(t, "z", r.Alias)
	assert.Equal(t, "-", r.AliasIndicator)
	assert.Equal(t, "td>-attributes", r.Parameters[0])
	assert.Equal(t, "some description", r.Description)
}
