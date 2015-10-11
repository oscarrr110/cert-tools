package cmd


import (
	"fmt"
	"os"
	"github.com/oscarrr110/cert-tools/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/oscarrr110/cert-tools/utils"
)


const (
	caPrefix  = "ca"
	caKeyPrefix  = "key"
	crtSuffix     = ".pem"
)

const (
	rootPerm   = 0400
	branchPerm = 0440
	leafPerm   = 0444
)

func NewInitCommand() cli.Command {
	return cli.Command{
		Name:        "init",
		Usage:       "Create Certificate Authority",
		Description: "Create Certificate Authority, including certificate and key pem file.",
		Flags: []cli.Flag{
			cli.StringFlag{"ip", "", "IP address of the host", ""},
			cli.StringFlag{"domain", "", "Use domain instead of IP address for SAN", ""},
			cli.IntFlag{"key-bits", 2048, "Bit size of RSA keypair to generate", ""},
			cli.IntFlag{"years", 10, "How long until the CA certificate expires", ""},
			cli.StringFlag{"organization", "cos-ca", "CA Certificate organization", ""},
		},
		Action: initAction,
	}
}

func initAction(c *cli.Context) {

	err := utils.GenerateCACertificate(d.Path(caPrefix+crtSuffix), d.Path(caKeyPrefix+crtSuffix), c.String("organization"), c.Int("years"), c.Int("key-bits"))
//	key, err := pkix.CreateRSAKey(c.Int("key-bits"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Create CA cert error:", err)
		os.Exit(1)
	} else {
		fmt.Println("Created ca/key complete")
	}
}