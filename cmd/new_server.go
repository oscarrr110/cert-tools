package cmd

import (
	"fmt"
	"os"
	"github.com/oscarrr110/cert-tools/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/oscarrr110/cert-tools/utils"
	"strings"
)

func NewNewCertCommand() cli.Command {
	return cli.Command{
		Name:        "new-cert",
		Usage:       "Create certificate request for host",
		Description: "Create certificate for host, including certificate signing request and key. Certificate could be generated by signing the request.",
		Flags: []cli.Flag{
			cli.StringFlag{"ips", "127.0.0.1", "IP address of the host", ""},
			cli.IntFlag{"key-bits", 2048, "Bit size of RSA keypair to generate", ""},
			cli.StringFlag{"organization", "cos-service", "CA Certificate organization", ""},
			cli.StringFlag{"common-name", "user|role|path", "Certificate common name", ""},
			cli.IntFlag{"years", 10, "How long until the CA certificate expires", ""},
			cli.StringFlag{"country", "USA", "Certificate country", ""},
		},
		Action: newCertAction,
	}
}

func newCertAction(c *cli.Context) {

	if len(c.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "cert file must be provided.")
		os.Exit(1)
	}

	certFileName := c.Args()[0]
	years := c.Int("years")
	bits := c.Int("key-bits")

	if c.IsSet("passphrase") {

	}

	ips := c.String("ips")
	ipArray := strings.Split(ips, ",")

	certFile := certFileName + crtSuffix
	keyFile := certFileName + "-key" + crtSuffix
	caFile := caPrefix + crtSuffix
	caKeyFile := caKeyPrefix + crtSuffix
	org := c.String("organization")
	commonName := c.String("common-name")


	err := utils.GenerateCert(ipArray, d.Path(certFile), d.Path(keyFile), d.Path(caFile), d.Path(caKeyFile), org, commonName, years,bits)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Create Servcer cert error:", err)
		os.Exit(1)
	} else {
		fmt.Println("Create server/key complete")
	}

}