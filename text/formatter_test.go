package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDescriptionParsing(t *testing.T) {
	str := `
See also \fI--dns-interface\fP and \fI--dns-ipv4-addr\fP. \fI--dns-ipv6-addr\fP \fI--dns-ipv6-addr\fP requires that the underlying libcurl was built to support c-ares.
Added in 7.33.0.`

	expected := `
See also ["0"][#fff000]--dns-interface[white][""] and ["1"][#fff000]--dns-ipv4-addr[white][""]. ["2"][#fff000]--dns-ipv6-addr[white][""] ["3"][#fff000]--dns-ipv6-addr[white][""] requires that the underlying libcurl was built to support c-ares.
Added in 7.33.0.`

	str = FormatDescription(str)

	assert.Equal(t, expected, str)
}

func TestStripRegions(t *testing.T) {
	str := `[#fff000]--dns-interface[white]`

	expected := `--dns-interface`

	str = StripColor(str)

	assert.Equal(t, expected, str)
}
