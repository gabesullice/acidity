package main

import (
	"fmt"
	"os"

	k "github.com/alecthomas/kingpin"
)

const (
	DEFAULT_ENDPOINT string = "https://www.acidity.io"
)

var (
	app      = k.New("acidity", "Command line ACI tool.")
	TLSCert  = app.Flag("cert-file", "TLS Certificate file.").Envar("ACIDITY_CERT_FILE").File()
	TLSKey   = app.Flag("key-file", "TLS Client key file.").Envar("ACIDITY_KEY_FILE").File()
	TLSCA    = app.Flag("ca-file", "CA bundle file to verify peers.").Envar("ACIDITY_CA_FILE").File()
	Endpoint = app.Flag("endpoint", "An endpoint to use for publishing your ACI.").Envar("ACIDITY_ENDPOINT").Default(DEFAULT_ENDPOINT).URL()

	publish    = app.Command("publish", "Publish an ACI")
	publishACI = publish.Arg("aci", "The file path to the ACI to publish.").File()
)

func main() {
	switch k.MustParse(app.Parse(os.Args[1:])) {
	case publish.FullCommand():
		aci := *publishACI
		url := *Endpoint
		fmt.Printf("Publish to %s to %s\n", aci.Name(), url.String())
	}
}
