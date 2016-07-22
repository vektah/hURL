package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"

	"github.com/99designs/httpsignatures-go"
	"github.com/alecthomas/kingpin"
	"github.com/vektah/hURL/lib"
)

func main() {
	verbose := kingpin.Flag("verbose", "Make the operation more talkative").Short('v').Bool()
	sigId := kingpin.Flag("sig-id", "The http-signatures id").String()
	sigKey := kingpin.Flag("sig-key", "The http-signatures key").String()
	insecure := kingpin.Flag("insecure", "Disable TLS cert verification").Bool()

	url := kingpin.Arg("url", "The url to fetch").Required().String()

	kingpin.Parse()

	r, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		kingpin.FatalUsage(err.Error())
	}

	if *sigId != "" && *sigKey != "" {
		httpsignatures.DefaultSha256Signer.SignRequest(*sigId, *sigKey, r)
	}

	if *verbose {
		r.Write(lib.PrefixLines(os.Stdout, "> "))
		println()
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: *insecure},
		},
	}

	resp, err := client.Do(r)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	if *verbose {
		resp.Header.Write(lib.PrefixLines(os.Stdout, "< "))
		println()
	}

	io.Copy(os.Stdout, resp.Body)

	if resp.StatusCode != 200 {
		os.Exit(1)
	}
}
