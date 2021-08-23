package man

import (
	"github.com/ignasbernotas/explain/parsers/man/data"
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
	result := parser.Parse(TestOptions, "curl")

	opts := result.Options.Options()

	assert.Equal(t, 3, len(opts))
	assert.Equal(t, "cacert", opts[0].Name)
	assert.Equal(t, "cacert", opts[0].Alias)
	assert.Equal(t, "append", opts[1].Name)
	assert.Equal(t, "a", opts[1].Alias)
	assert.Equal(t, "basic", opts[2].Name)
	assert.Equal(t, "basic", opts[2].Alias)
}

func TestSectionLineParsing(t *testing.T) {
	parser := NewParser()
	result := parser.parseSectionLine(`.SH DESCRIPTION`)
	assert.Equal(t, SectionDescription, result)

	result = parser.parseSectionLine(`.SH "OPTIONS"`)
	assert.Equal(t, SectionOptions, result)
}

func TestOptionLineParsing(t *testing.T) {
	parser := NewParser()
	result := parser.parseOptionLine(`.IP "--alt-svc <file name>"`, "")
	assert.Equal(t, `--alt-svc <file name>`, result)

	result = parser.parseOptionLine(`.It Fl 4 adadsas`, "")
	assert.Equal(t, `4`, result)
}

func TestOptionLineParsing2(t *testing.T) {
	parser := NewParser()
	parser.Parse(data.SSH_OUTPUT, "ssh")
	//assert.Equal(t, `4`, result)
}

func TestOptionLineParsing3(t *testing.T) {
	parser := NewParser()
	parser.Parse(data.AB_OUTPUT, "ab")
	//assert.Equal(t, `4`, result.)
}
