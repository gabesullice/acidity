package main

import (
	"fmt"
	"os"

	k "github.com/alecthomas/kingpin"
)

const (
	DEFAULT_ENDPOINT    string = "https://www.acidity.io"
	DEFAULT_AUTH_SERVER string = "https://auth.acidity.io"
)

var (
	app      = k.New("acidity", "Command line ACI tool.")
	TLSCert  = app.Flag("cert-file", "TLS Certificate file.").Envar("ACIDITY_CERT_FILE").File()
	TLSKey   = app.Flag("key-file", "TLS Client key file.").Envar("ACIDITY_KEY_FILE").File()
	TLSCA    = app.Flag("ca-file", "CA bundle file to verify peers.").Envar("ACIDITY_CA_FILE").File()
	Endpoint = app.Flag("endpoint", "An endpoint to use for publishing and discovering ACIs.").Short('e').Envar("ACIDITY_ENDPOINT").Default(DEFAULT_ENDPOINT).URL()

	publish              = app.Command("publish", "Publish an ACI")
	publishACI           = publish.Arg("aci", "The file path to the ACI to publish.").Required().File()
	publishPrivate       = publish.Flag("private", "Denotes that the ACI should be published to the peer network privately.").Bool()
	publishPrivateCA     = publish.Flag("private-ca", "CA file to publish with the ACI.").Envar("ACIDITY_PRIVATE_CA").File()
	publishAuthenticated = publish.Flag("authenticated", "Publishes the ACI with a third party certificate.").Bool()
	publishAuthServer    = publish.Flag("auth-server", "An authentication server.").Envar("ACIDITY_AUTH_SERVER").Default(DEFAULT_AUTH_SERVER).URL()

	fetch       = app.Command("fetch", "Download an ACI from an acidity peer network.")
	fetchMagnet = fetch.Flag("magnet", "Magnet link from which to obtain the ACI. If unspecified, an attempt will be made to find one.").Short('m').URL()
	fetchACI    = fetch.Arg("name", "The name of the ACI to retrieve.").Required().String()
)

func main() {
	switch k.MustParse(app.Parse(os.Args[1:])) {
	case publish.FullCommand():
		validatePublish()
		aci := *publishACI
		url := *Endpoint
		fmt.Printf("Publishing %s to peers from %s...\n", aci.Name(), url.String())
		break
	case fetch.FullCommand():
		aci := *fetchACI
		fmt.Printf("Fetching %s from peers...\n", aci)
	}
}

func validatePublish() {
	if (*publishPrivate || *publishAuthenticated) && *publishPrivate == *publishAuthenticated {
		k.FatalUsage("An ACI may not be published as private AND authenticated.\n")
	}
	if *publishPrivate && *publishPrivateCA == nil {
		k.FatalUsage("If you publish an ACI privately, you must provide a private CA file.\n")
	}
}
