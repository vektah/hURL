hURL
====

Simple command line tool for hitting http-signatures based apis


### Installation
```sh
go get -v github.com/vektah/hurl
```

### Usage
```
usage: hurl [<flags>] <url>

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose          Make the operation more talkative
      --sig-id=SIG-ID    The http-signatures id
      --sig-key=SIG-KEY  The http-signatures key
      --insecure         Disable TLS cert verification

Args:
  <url>  The url to fetch
```

A basic example
```sh
hurl https://example.org/api/bob --sig-id MyId--sig-key MyKey --insecure
```
