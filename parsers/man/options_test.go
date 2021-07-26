package man

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOption(t *testing.T) {
	r := buildOption("-m, --max-time <seconds>", "some description")

	assert.Equal(t, "max-time", r.Name)
	assert.Equal(t, "--", r.NameIndicator)
	assert.Equal(t, "m", r.ShortName)
	assert.Equal(t, "-", r.ShortNameIndicator)
	assert.Equal(t, "seconds", r.Parameter)
	assert.Equal(t, "some description", r.Description)
}
