package man

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const TestOptions = `
.SH OPTIONS
.IP "--cacert <file>"
(TLS) Tells curl to use the specified certificate file to verify the peer. The file
may contain multiple CA certificates. The certificate(s) must be in PEM
format. Normally curl is built to use a default file for this, so this option
is typically used to alter that default file.
.IP "-a, --append"
(FTP SFTP) When used in an upload, this makes curl append to the target file instead of
overwriting it. If the remote file doesn't exist, it will be created.  Note
that this flag is ignored by some SFTP servers (including OpenSSH).
.IP "--basic"
(HTTP) Tells curl to use HTTP Basic authentication with the remote host. This is the
default and this option is usually pointless, unless you use it to override a
previously set option that sets a different authentication method (such as
\fI--ntlm\fP, \fI--digest\fP, or \fI--negotiate\fP).
.IP "--empty"
`

func TestParsing(t *testing.T) {
	parser := NewParser()
	result := parser.Parse(TestOptions)

	assert.Equal(t, 3, len(result.Options))
	assert.Equal(t, "cacert", result.Options[0].Name)
	assert.Equal(t, "cacert", result.Options[0].ShortName)
	assert.Equal(t, "append", result.Options[1].Name)
	assert.Equal(t, "a", result.Options[1].ShortName)
	assert.Equal(t, "basic", result.Options[2].Name)
	assert.Equal(t, "basic", result.Options[2].ShortName)
}

func TestSectionLineParsing(t *testing.T) {
	parser := NewParser()
	result := parser.parseSectionLine(`.SH DESCRIPTION`)
	assert.Equal(t, SectionDescription, result)

	result = parser.parseSectionLine(`.SH OPTIONS`)
	assert.Equal(t, SectionOptions, result)
}

func TestOptionLineParsing(t *testing.T) {
	parser := NewParser()
	result := parser.parseOptionLine(`.IP "--alt-svc <file name>"`)
	assert.Equal(t, `--alt-svc <file name>`, result)

	result = parser.parseOptionLine(`.It Fl 4`)
	assert.Equal(t, `4`, result)
}
