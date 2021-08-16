package man

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

const sentenceText = `Connect to the target host by first making a
.Nm
connection to the jump host described by
.Ar destination
and then establishing a TCP forwarding to the ultimate destination from
there.
Multiple jump hops may be specified separated by comma characters.
This is a shortcut to specify a
.Cm ProxyJump
configuration directive.
Note that configuration directives supplied on the command-line generally
apply to the destination host and not any specified jump hosts.
Use
.Pa ~/.ssh/config
to specify configuration for jump hosts.`

const sentenceText2 = `Enables forwarding of the authentication agent connection.
This can also be specified on a per-host basis in a configuration file.

Agent forwarding should be enabled with caution.
Users with the ability to bypass file permissions on the remote host
(for the agent's
.Ux Ns -domain
socket) can access the local agent through the forwarded connection.
An attacker cannot obtain key material from the agent,
however they can perform operations on the keys that enable them to
authenticate using the identities loaded into the agent.
.Pp
`

const sentenceText3 = `.It Fl L Xo
.Sm off
.Oo Ar bind_address : Oc
.Ar port : host : hostport
.Sm on
.Xc
.It Fl L Xo
.Sm off
.Oo Ar bind_address : Oc
.Ar port : remote_socket
.Sm on
.Xc
.It Fl L Xo
.Sm off
.Ar local_socket : host : hostport
.Sm on
.Xc
.It Fl L Xo
.Sm off
.Ar local_socket : remote_socket
.Sm on
.Xc
Specifies that connections to the given TCP port or Unix socket on the local
(client) host are to be forwarded to the given host and port, or Unix socket,
on the remote side.
This works by allocating a socket to listen to either a TCP
.Ar port
on the local side, optionally bound to the specified
.Ar bind_address ,
or to a Unix socket.
Whenever a connection is made to the local port or socket, the
connection is forwarded over the secure channel, and a connection is
made to either
.Ar host
port
.Ar hostport ,
or the Unix socket
.Ar remote_socket ,
from the remote machine.
.Pp
Port forwardings can also be specified in the configuration file.
Only the superuser can forward privileged ports.
IPv6 addresses can be specified by enclosing the address in square brackets.
.Pp
By default, the local port is bound in accordance with the
.Cm GatewayPorts
setting.
However, an explicit
.Ar bind_address
may be used to bind the connection to a specific address.
The
.Ar bind_address
of
.Dq localhost
indicates that the listening port be bound for local use only, while an
empty address or
.Sq *
indicates that the port should be available from all interfaces.
.Pp`

// remove all newlines where .\n
func TestSentences(t *testing.T) {
	spew.Dump(sentenceText3)
	s := wrapReplace(sentenceText3)
	spew.Dump(s)
	s = replaceTokens(s)
	spew.Dump(s)
	a := fixSentences(s)
	spew.Dump(a)
	//assert.Equal(t, "basic", opts[2].Alias)
}

func TestWrap(t *testing.T) {
	str := ".Xr gzip 1 ."

	str = wrapReference(str)

	spew.Dump(str)
}

func TestWrap2(t *testing.T) {
	s := "By default, the local port is bound in accordance with the\n.Cm GatewayPorts\nsetting."
	spew.Dump(s)
	s = wrapReference(s)
	spew.Dump(s)
	s = wrapReplace(s)
	spew.Dump(s)
	s = replaceTokens(s)
	spew.Dump(s)
	a := fixSentences(s)
	spew.Dump(a)
}
func TestWrap3(t *testing.T) {
	s := ".It Fl A\nEnables forwarding of the authentication agent connection.\nThis can also be specified on a per-host basis in a configuration file.\n.Pp\nAgent forwarding should be enabled with caution.\nUsers with the ability to bypass file permissions on the remote host\n(for the agent's\n.Ux Ns -domain\nsocket) can access the local agent through the forwarded connection.\nAn attacker cannot obtain key material from the agent,\nhowever they can perform operations on the keys that enable them to\nauthenticate using the identities loaded into the agent.\n.Pp"
	spew.Dump(s)
	s = wrapReference(s)
	spew.Dump(s)
	s = wrapReplace(s)
	spew.Dump(s)
	s = replaceTokens(s)
	spew.Dump(s)
	a := fixSentences(s)
	spew.Dump(a)
}