package data

const CURL_OUTPUT = `
.SH OPTIONS
.IP "--abstract-unix-socket <path>"
(HTTP) Connect through an abstract Unix domain socket, instead of using the network.
Note: netstat shows the path of an abstract socket prefixed with '@', however
the <path> argument should not have this leading character.
.IP "--alt-svc <file name>"
(HTTPS) WARNING: this option is experiemental. Do not use in production.

This option enables the alt-svc parser in curl. If the file name points to an
existing alt-svc cache file, that will be used. After a completed transfer,
the cache will be saved to the file name again if it has been modified.

Specifiy a "" file name (zero length) to avoid loading/saving and make curl
just handle the cache in memory.

If this option is used several times, curl will load contents from all the
files but the the last one will be used for saving.

Added in 7.64.1.
.IP "--anyauth"
(HTTP) Tells curl to figure out authentication method by itself, and use the most
secure one the remote site claims to support. This is done by first doing a
request and checking the response-headers, thus possibly inducing an extra
network round-trip. This is used instead of setting a specific authentication
method, which you can do with \fI--basic\fP, \fI--digest\fP, \fI--ntlm\fP, and \fI--negotiate\fP.

Using \fI--anyauth\fP is not recommended if you do uploads from stdin, since it may
require data to be sent twice and then the client must be able to rewind. If
the need should arise when uploading from stdin, the upload operation will
fail.

Used together with \fI-u, --user\fP.

See also \fI--proxy-anyauth\fP and \fI--basic\fP and \fI--digest\fP.
.IP "-a, --append"
(FTP SFTP) When used in an upload, this makes curl append to the target file instead of
overwriting it. If the remote file doesn't exist, it will be created.  Note
that this flag is ignored by some SFTP servers (including OpenSSH).
.IP "--basic"
(HTTP) Tells curl to use HTTP Basic authentication with the remote host. This is the
default and this option is usually pointless, unless you use it to override a
previously set option that sets a different authentication method (such as
\fI--ntlm\fP, \fI--digest\fP, or \fI--negotiate\fP).

Used together with \fI-u, --user\fP.

See also \fI--proxy-basic\fP.
.IP "--cacert <file>"
(TLS) Tells curl to use the specified certificate file to verify the peer. The file
may contain multiple CA certificates. The certificate(s) must be in PEM
format. Normally curl is built to use a default file for this, so this option
is typically used to alter that default file.

curl recognizes the environment variable named 'CURL_CA_BUNDLE' if it is
set, and uses the given path as a path to a CA cert bundle. This option
overrides that variable.

The windows version of curl will automatically look for a CA certs file named
\'curl-ca-bundle.crt\', either in the same directory as curl.exe, or in the
Current Working Directory, or in any folder along your PATH.

If curl is built against the NSS SSL library, the NSS PEM PKCS#11 module
(libnsspem.so) needs to be available for this option to work properly.

(iOS and macOS only) If curl is built against Secure Transport, then this
option is supported for backward compatibility with other SSL engines, but it
should not be set. If the option is not set, then curl will use the
certificates in the system and user Keychain to verify the peer, which is the
preferred method of verifying the peer's certificate chain.

(Schannel only) This option is supported for Schannel in Windows 7 or later with
libcurl 7.60 or later. This option is supported for backward compatibility
with other SSL engines; instead it is recommended to use Windows' store of
root certificates (the default for Schannel).

If this option is used several times, the last one will be used.
.IP "--capath <dir>"
(TLS) Tells curl to use the specified certificate directory to verify the
peer. Multiple paths can be provided by separating them with ":" (e.g.
\&"path1:path2:path3"). The certificates must be in PEM format, and if curl is
built against OpenSSL, the directory must have been processed using the
c_rehash utility supplied with OpenSSL. Using \fI--capath\fP can allow
OpenSSL-powered curl to make SSL-connections much more efficiently than using
\fI--cacert\fP if the --cacert file contains many CA certificates.

If this option is set, the default capath value will be ignored, and if it is
used several times, the last one will be used.
.IP "--cert-status"
(TLS) Tells curl to verify the status of the server certificate by using the
Certificate Status Request (aka. OCSP stapling) TLS extension.

If this option is enabled and the server sends an invalid (e.g. expired)
response, if the response suggests that the server certificate has been revoked,
or no response at all is received, the verification fails.

This is currently only implemented in the OpenSSL, GnuTLS and NSS backends.

Added in 7.41.0.
.IP "--cert-type <type>"
(TLS) Tells curl what type the provided client certificate is using. PEM, DER, ENG
and P12 are recognized types.  If not specified, PEM is assumed.

If this option is used several times, the last one will be used.

See also \fI-E, --cert\fP and \fI--key\fP and \fI--key-type\fP.
.IP "-E, --cert <certificate[:password]>"
(TLS) Tells curl to use the specified client certificate file when getting a file
with HTTPS, FTPS or another SSL-based protocol. The certificate must be in
PKCS#12 format if using Secure Transport, or PEM format if using any other
engine.  If the optional password isn't specified, it will be queried for on
the terminal. Note that this option assumes a \&"certificate" file that is the
private key and the client certificate concatenated! See \fI-E, --cert\fP and \fI--key\fP to
specify them independently.

If curl is built against the NSS SSL library then this option can tell
curl the nickname of the certificate to use within the NSS database defined
by the environment variable SSL_DIR (or by default /etc/pki/nssdb). If the
NSS PEM PKCS#11 module (libnsspem.so) is available then PEM files may be
loaded. If you want to use a file from the current directory, please precede
it with "./" prefix, in order to avoid confusion with a nickname.  If the
nickname contains ":", it needs to be preceded by "\\" so that it is not
recognized as password delimiter.  If the nickname contains "\\", it needs to
be escaped as "\\\\" so that it is not recognized as an escape character.

If curl is built against OpenSSL library, and the engine pkcs11 is available,
then a PKCS#11 URI (RFC 7512) can be used to specify a certificate located in
a PKCS#11 device. A string beginning with "pkcs11:" will be interpreted as a
PKCS#11 URI. If a PKCS#11 URI is provided, then the \fI--engine\fP option will be set
as "pkcs11" if none was provided and the \fI--cert-type\fP option will be set as
"ENG" if none was provided.

(iOS and macOS only) If curl is built against Secure Transport, then the
certificate string can either be the name or public key hash of a certificate/private
key in the system or user keychain, or the path to a PKCS#12-encoded certificate and
private key. A string beginning with "pkh/" will be interpreted as a
public key hash. If you want to use a file from the current directory, please
precede it with "./" prefix, in order to avoid confusion with a nickname.

(Schannel only) Client certificates must be specified by a path
expression to a certificate store. (Loading PFX is not supported; you can
import it to a store first). You can use
"<store location>\\<store name>\\<thumbprint>" to refer to a certificate
in the system certificates store, for example,
"CurrentUser\\MY\\934a7ac6f8a5d579285a74fa61e19f23ddfe8d7a". Thumbprint is
usually a SHA-1 hex string which you can see in certificate details. Following
store locations are supported: CurrentUser, LocalMachine, CurrentService,
Services, CurrentUserGroupPolicy, LocalMachineGroupPolicy,
LocalMachineEnterprise.

If this option is used several times, the last one will be used.

See also \fI--cert-type\fP and \fI--key\fP and \fI--key-type\fP.
.IP "--ciphers <list of ciphers>"
(TLS) Specifies which ciphers to use in the connection. The list of ciphers must
specify valid ciphers. Read up on SSL cipher list details on this URL:

 https://curl.haxx.se/docs/ssl-ciphers.html

If this option is used several times, the last one will be used.
.IP "--compressed-ssh"
(SCP SFTP) Enables built-in SSH compression.
This is a request, not an order; the server may or may not do it.

Added in 7.56.0.
.IP "--compressed"
(HTTP) Request a compressed response using one of the algorithms curl supports, and
save the uncompressed document.  If this option is used and the server sends
an unsupported encoding, curl will report an error.
.IP "-K, --config <file>"

Specify a text file to read curl arguments from. The command line arguments
found in the text file will be used as if they were provided on the command
line.

Options and their parameters must be specified on the same line in the file,
separated by whitespace, colon, or the equals sign. Long option names can
optionally be given in the config file without the initial double dashes and
if so, the colon or equals characters can be used as separators. If the option
is specified with one or two dashes, there can be no colon or equals character
between the option and its parameter.

If the parameter is to contain whitespace, the parameter must be enclosed
within quotes. Within double quotes, the following escape sequences are
available: \\\\, \\", \\t, \\n, \\r and \\v. A backslash preceding any other
letter is ignored. If the first column of a config line is a '#' character,
the rest of the line will be treated as a comment. Only write one option per
physical line in the config file.

Specify the filename to \fI-K, --config\fP as '-' to make curl read the file from stdin.

Note that to be able to specify a URL in the config file, you need to specify
it using the \fI--url\fP option, and not by simply writing the URL on its own
line. So, it could look similar to this:

url = "https://curl.haxx.se/docs/"

When curl is invoked, it (unless \fI-q, --disable\fP is used) checks for a default
config file and uses it if found. The default config file is checked for in
the following places in this order:

1) curl tries to find the "home dir": It first checks for the CURL_HOME and
then the HOME environment variables. Failing that, it uses getpwuid() on
Unix-like systems (which returns the home dir given the current user in your
system). On Windows, it then checks for the APPDATA variable, or as a last
resort the '%USERPROFILE%\\Application Data'.

2) On windows, if there is no _curlrc file in the home dir, it checks for one
in the same dir the curl executable is placed. On Unix-like systems, it will
simply try to load .curlrc from the determined home dir.

.nf
# --- Example file ---
# this is a comment
url = "example.com"
output = "curlhere.html"
user-agent = "superagent/1.0"

# and fetch another URL too
url = "example.com/docs/manpage.html"
-O
referer = "http://nowhereatall.example.com/"
# --- End of example file ---
.fi

This option can be used multiple times to load multiple config files.
.IP "--connect-timeout <seconds>"
Maximum time in seconds that you allow curl's connection to take.  This only
limits the connection phase, so if curl connects within the given period it
will continue - if not it will exit.  Since version 7.32.0, this option
accepts decimal values.

If this option is used several times, the last one will be used.

See also \fI-m, --max-time\fP.
.IP "--connect-to <HOST1:PORT1:HOST2:PORT2>"

For a request to the given HOST1:PORT1 pair, connect to HOST2:PORT2 instead.
This option is suitable to direct requests at a specific server, e.g. at a
specific cluster node in a cluster of servers. This option is only used to
establish the network connection. It does NOT affect the hostname/port that is
used for TLS/SSL (e.g. SNI, certificate verification) or for the application
protocols. "HOST1" and "PORT1" may be the empty string, meaning "any
host/port". "HOST2" and "PORT2" may also be the empty string, meaning "use the
request's original host/port".

A "host" specified to this option is compared as a string, so it needs to
match the name used in request URL. It can be either numerical such as
"127.0.0.1" or the full host name such as "example.org".

This option can be used many times to add many connect rules.

See also \fI--resolve\fP and \fI-H, --header\fP. Added in 7.49.0.
.IP "-C, --continue-at <offset>"
Continue/Resume a previous file transfer at the given offset. The given offset
is the exact number of bytes that will be skipped, counting from the beginning
of the source file before it is transferred to the destination.  If used with
uploads, the FTP server command SIZE will not be used by curl.

Use "-C -" to tell curl to automatically find out where/how to resume the
transfer. It then uses the given output/input files to figure that out.

If this option is used several times, the last one will be used.

See also \fI-r, --range\fP.
.IP "-c, --cookie-jar <filename>"
(HTTP) Specify to which file you want curl to write all cookies after a completed
operation. Curl writes all cookies from its in-memory cookie storage to the
given file at the end of operations. If no cookies are known, no data will be
written. The file will be written using the Netscape cookie file format. If
you set the file name to a single dash, "-", the cookies will be written to
stdout.

This command line option will activate the cookie engine that makes curl
record and use cookies. Another way to activate it is to use the \fI-b, --cookie\fP
option.

If the cookie jar can't be created or written to, the whole curl operation
won't fail or even report an error clearly. Using \fI-v, --verbose\fP will get a warning
displayed, but that is the only visible feedback you get about this possibly
lethal situation.

If this option is used several times, the last specified file name will be
used.
.IP "-b, --cookie <data|filename>"
(HTTP) Pass the data to the HTTP server in the Cookie header. It is supposedly
the data previously received from the server in a "Set-Cookie:" line.  The
data should be in the format "NAME1=VALUE1; NAME2=VALUE2".

If no '=' symbol is used in the argument, it is instead treated as a filename
to read previously stored cookie from. This option also activates the cookie
engine which will make curl record incoming cookies, which may be handy if
you're using this in combination with the \fI-L, --location\fP option or do multiple URL
transfers on the same invoke. If the file name is exactly a minus ("-"), curl
will instead the contents from stdin.

The file format of the file to read cookies from should be plain HTTP headers
(Set-Cookie style) or the Netscape/Mozilla cookie file format.

The file specified with \fI-b, --cookie\fP is only used as input. No cookies will be
written to the file. To store cookies, use the \fI-c, --cookie-jar\fP option.

Exercise caution if you are using this option and multiple transfers may
occur.  If you use the NAME1=VALUE1; format, or in a file use the Set-Cookie
format and don't specify a domain, then the cookie is sent for any domain
(even after redirects are followed) and cannot be modified by a server-set
cookie. If the cookie engine is enabled and a server sets a cookie of the same
name then both will be sent on a future transfer to that server, likely not
what you intended.  To address these issues set a domain in Set-Cookie (doing
that will include sub domains) or use the Netscape format.

If this option is used several times, the last one will be used.

Users very often want to both read cookies from a file and write updated
cookies back to a file, so using both \fI-b, --cookie\fP and \fI-c, --cookie-jar\fP in the same
command line is common.
.IP "--create-dirs"
When used in conjunction with the \fI-o, --output\fP option, curl will create the
necessary local directory hierarchy as needed. This option creates the dirs
mentioned with the \fI-o, --output\fP option, nothing else. If the --output file name
uses no dir or if the dirs it mentions already exist, no dir will be created.

To create remote directories when using FTP or SFTP, try \fI--ftp-create-dirs\fP.
.IP "--crlf"
(FTP SMTP) Convert LF to CRLF in upload. Useful for MVS (OS/390).

(SMTP added in 7.40.0)
.IP "--crlfile <file>"
(TLS) Provide a file using PEM format with a Certificate Revocation List that may
specify peer certificates that are to be considered revoked.

If this option is used several times, the last one will be used.

Added in 7.19.7.
.IP "--data-ascii <data>"
(HTTP) This is just an alias for \fI-d, --data\fP.
.IP "--data-binary <data>"
(HTTP) This posts data exactly as specified with no extra processing whatsoever.

If you start the data with the letter @, the rest should be a filename.  Data
is posted in a similar manner as \fI-d, --data\fP does, except that newlines and
carriage returns are preserved and conversions are never done.

Like \fI-d, --data\fP the default content-type sent to the server is
application/x-www-form-urlencoded. If you want the data to be treated as
arbitrary binary data by the server then set the content-type to octet-stream:
-H "Content-Type: application/octet-stream".

If this option is used several times, the ones following the first will append
data as described in \fI-d, --data\fP.
.IP "--data-raw <data>"
(HTTP) This posts data similarly to \fI-d, --data\fP but without the special
interpretation of the @ character.

See also \fI-d, --data\fP. Added in 7.43.0.
.IP "--data-urlencode <data>"
(HTTP) This posts data, similar to the other \fI-d, --data\fP options with the exception
that this performs URL-encoding.

To be CGI-compliant, the <data> part should begin with a \fIname\fP followed
by a separator and a content specification. The <data> part can be passed to
curl using one of the following syntaxes:
.RS
.IP "content"
This will make curl URL-encode the content and pass that on. Just be careful
so that the content doesn't contain any = or @ symbols, as that will then make
the syntax match one of the other cases below!
.IP "=content"
This will make curl URL-encode the content and pass that on. The preceding =
symbol is not included in the data.
.IP "name=content"
This will make curl URL-encode the content part and pass that on. Note that
the name part is expected to be URL-encoded already.
.IP "@filename"
This will make curl load data from the given file (including any newlines),
URL-encode that data and pass it on in the POST.
.IP "name@filename"
This will make curl load data from the given file (including any newlines),
URL-encode that data and pass it on in the POST. The name part gets an equal
sign appended, resulting in \fIname=urlencoded-file-content\fP. Note that the
name is expected to be URL-encoded already.
.RE

See also \fI-d, --data\fP and \fI--data-raw\fP. Added in 7.18.0.
.IP "-d, --data <data>"
(HTTP) Sends the specified data in a POST request to the HTTP server, in the same way
that a browser does when a user has filled in an HTML form and presses the
submit button. This will cause curl to pass the data to the server using the
content-type application/x-www-form-urlencoded.  Compare to \fI-F, --form\fP.

\fI--data-raw\fP is almost the same but does not have a special interpretation of
the @ character. To post data purely binary, you should instead use the
\fI--data-binary\fP option.  To URL-encode the value of a form field you may use
\fI--data-urlencode\fP.

If any of these options is used more than once on the same command line, the
data pieces specified will be merged together with a separating
&-symbol. Thus, using '-d name=daniel -d skill=lousy' would generate a post
chunk that looks like \&'name=daniel&skill=lousy'.

If you start the data with the letter @, the rest should be a file name to
read the data from, or - if you want curl to read the data from
stdin. Multiple files can also be specified. Posting data from a file named
'foobar' would thus be done with \fI-d, --data\fP @foobar. When --data is told to read
from a file like that, carriage returns and newlines will be stripped out. If
you don't want the @ character to have a special interpretation use \fI--data-raw\fP
instead.

See also \fI--data-binary\fP and \fI--data-urlencode\fP and \fI--data-raw\fP. This option overrides \fI-F, --form\fP and \fI-I, --head\fP and \fI-T, --upload-file\fP.
.IP "--delegation <LEVEL>"
(GSS/kerberos) Set LEVEL to tell the server what it is allowed to delegate when it
comes to user credentials.
.RS
.IP "none"
Don't allow any delegation.
.IP "policy"
Delegates if and only if the OK-AS-DELEGATE flag is set in the Kerberos
service ticket, which is a matter of realm policy.
.IP "always"
Unconditionally allow the server to delegate.
.RE
.IP "--digest"
(HTTP) Enables HTTP Digest authentication. This is an authentication scheme that
prevents the password from being sent over the wire in clear text. Use this in
combination with the normal \fI-u, --user\fP option to set user name and password.

If this option is used several times, only the first one is used.

See also \fI-u, --user\fP and \fI--proxy-digest\fP and \fI--anyauth\fP. This option overrides \fI--basic\fP and \fI--ntlm\fP and \fI--negotiate\fP.
.IP "--disable-eprt"
(FTP) Tell curl to disable the use of the EPRT and LPRT commands when doing active
FTP transfers. Curl will normally always first attempt to use EPRT, then LPRT
before using PORT, but with this option, it will use PORT right away. EPRT and
LPRT are extensions to the original FTP protocol, and may not work on all
servers, but they enable more functionality in a better way than the
traditional PORT command.

--eprt can be used to explicitly enable EPRT again and --no-eprt is an alias
for \fI--disable-eprt\fP.

If the server is accessed using IPv6, this option will have no effect as EPRT
is necessary then.

Disabling EPRT only changes the active behavior. If you want to switch to
passive mode you need to not use \fI-P, --ftp-port\fP or force it with \fI--ftp-pasv\fP.
.IP "--disable-epsv"
(FTP) (FTP) Tell curl to disable the use of the EPSV command when doing passive FTP
transfers. Curl will normally always first attempt to use EPSV before PASV,
but with this option, it will not try using EPSV.

--epsv can be used to explicitly enable EPSV again and --no-epsv is an alias
for \fI--disable-epsv\fP.

If the server is an IPv6 host, this option will have no effect as EPSV is
necessary then.

Disabling EPSV only changes the passive behavior. If you want to switch to
active mode you need to use \fI-P, --ftp-port\fP.
.IP "-q, --disable"
If used as the first parameter on the command line, the \fIcurlrc\fP config
file will not be read and used. See the \fI-K, --config\fP for details on the default
config file search path.
.IP "--disallow-username-in-url"
(HTTP) This tells curl to exit if passed a url containing a username.

See also \fI--proto\fP. Added in 7.61.0.
.IP "--dns-interface <interface>"
(DNS) Tell curl to send outgoing DNS requests through <interface>. This option is a
counterpart to \fI--interface\fP (which does not affect DNS). The supplied string
must be an interface name (not an address).

See also \fI--dns-ipv4-addr\fP and \fI--dns-ipv6-addr\fP. \fI--dns-interface\fP requires that the underlying libcurl was built to support c-ares. Added in 7.33.0.
.IP "--dns-ipv4-addr <address>"
(DNS) Tell curl to bind to <ip-address> when making IPv4 DNS requests, so that
the DNS requests originate from this address. The argument should be a
single IPv4 address.

See also \fI--dns-interface\fP and \fI--dns-ipv6-addr\fP. \fI--dns-ipv4-addr\fP requires that the underlying libcurl was built to support c-ares. Added in 7.33.0.
.IP "--dns-ipv6-addr <address>"
(DNS) Tell curl to bind to <ip-address> when making IPv6 DNS requests, so that
the DNS requests originate from this address. The argument should be a
single IPv6 address.

See also \fI--dns-interface\fP and \fI--dns-ipv4-addr\fP. \fI--dns-ipv6-addr\fP requires that the underlying libcurl was built to support c-ares. Added in 7.33.0.
.IP "--dns-servers <addresses>"
Set the list of DNS servers to be used instead of the system default.
The list of IP addresses should be separated with commas. Port numbers
may also optionally be given as \fI:<port-number>\fP after each IP
address.

\fI--dns-servers\fP requires that the underlying libcurl was built to support c-ares. Added in 7.33.0.
.IP "--doh-url <URL>"
(all) Specifies which DNS-over-HTTPS (DOH) server to use to resolve hostnames,
instead of using the default name resolver mechanism. The URL must be HTTPS.

If this option is used several times, the last one will be used.
.IP "-D, --dump-header <filename>"
(HTTP FTP) Write the received protocol headers to the specified file.

This option is handy to use when you want to store the headers that an HTTP
site sends to you. Cookies from the headers could then be read in a second
curl invocation by using the \fI-b, --cookie\fP option! The \fI-c, --cookie-jar\fP option is a
better way to store cookies.

When used in FTP, the FTP server response lines are considered being "headers"
and thus are saved there.

If this option is used several times, the last one will be used.

See also \fI-o, --output\fP.
.IP "--egd-file <file>"
(TLS) Specify the path name to the Entropy Gathering Daemon socket. The socket is
used to seed the random engine for SSL connections.

See also \fI--random-file\fP.
.IP "--engine <name>"
(TLS) Select the OpenSSL crypto engine to use for cipher operations. Use \fI--engine\fP
list to print a list of build-time supported engines. Note that not all (or
none) of the engines may be available at run-time.
.IP "--expect100-timeout <seconds>"
(HTTP) Maximum time in seconds that you allow curl to wait for a 100-continue
response when curl emits an Expects: 100-continue header in its request. By
default curl will wait one second. This option accepts decimal values! When
curl stops waiting, it will continue as if the response has been received.

See also \fI--connect-timeout\fP. Added in 7.47.0.
.IP "--fail-early"
Fail and exit on the first detected transfer error.

When curl is used to do multiple transfers on the command line, it will
attempt to operate on each given URL, one by one. By default, it will ignore
errors if there are more URLs given and the last URL's success will determine
the error code curl returns. So early failures will be "hidden" by subsequent
successful transfers.

Using this option, curl will instead return an error on the first transfer
that fails, independent of the amount of URLs that are given on the command
line. This way, no transfer failures go undetected by scripts and similar.

This option is global and does not need to be specified for each use of \fI-:, --next\fP.

This option does not imply \fI-f, --fail\fP, which causes transfers to fail due to the
server's HTTP status code. You can combine the two options, however note \fI-f, --fail\fP
is not global and is therefore contained by \fI-:, --next\fP.

Added in 7.52.0.
.IP "-f, --fail"
(HTTP) Fail silently (no output at all) on server errors. This is mostly done to
better enable scripts etc to better deal with failed attempts. In normal cases
when an HTTP server fails to deliver a document, it returns an HTML document
stating so (which often also describes why and more). This flag will prevent
curl from outputting that and return error 22.

This method is not fail-safe and there are occasions where non-successful
response codes will slip through, especially when authentication is involved
(response codes 401 and 407).
.IP "--false-start"
(TLS) Tells curl to use false start during the TLS handshake. False start is a mode
where a TLS client will start sending application data before verifying the
server's Finished message, thus saving a round trip when performing a full
handshake.

This is currently only implemented in the NSS and Secure Transport (on iOS 7.0
or later, or OS X 10.9 or later) backends.

Added in 7.42.0.
.IP "--form-string <name=string>"
(HTTP SMTP IMAP) Similar to \fI-F, --form\fP except that the value string for the named parameter is used
literally. Leading \&'@' and \&'<' characters, and the \&';type=' string in
the value have no special meaning. Use this in preference to \fI-F, --form\fP if
there's any possibility that the string value may accidentally trigger the
\&'@' or \&'<' features of \fI-F, --form\fP.

See also \fI-F, --form\fP.
.IP "-F, --form <name=content>"
(HTTP SMTP IMAP) For HTTP protocol family, this lets curl emulate a filled-in form in which a
user has pressed the submit button. This causes curl to POST data using the
Content-Type multipart/form-data according to RFC 2388.

For SMTP and IMAP protocols, this is the mean to compose a multipart mail
message to transmit.
.SH FILES
.I ~/.curlrc
.RS
Default config file, see \fI-K, --config\fP for details.
