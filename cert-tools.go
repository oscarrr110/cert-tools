package main

import "fmt"
import "os"
import "github.com/oscarrr110/cert-tools/cmd"
import "github.com/oscarrr110/cert-tools/depot"
import "github.com/oscarrr110/cert-tools/Godeps/_workspace/src/github.com/codegangsta/cli"

func main() {

	//    err := utils.GenerateCACertificate("ca.pem", "key.pem","test-org",2048)
	//	     if err != nil {
	//		     fmt.Println("error: %v", err);
	//	     }

	app := cli.NewApp()
	app.Name = "cert-tools"
	app.Version = "0.0.1"
	app.Usage = "Simple tools for generating CA and its signed cert "
	app.Flags = []cli.Flag{
		cli.StringFlag{"depot-path", depot.DefaultFileDepotDir, "Location to store certificates, keys and other files.", ""},
	}

	app.Commands = []cli.Command{
		cmd.NewInitCommand(),
		cmd.NewNewCertCommand(),
	}

	app.Before = func(c *cli.Context) error {

		fmt.Println("Before running")
		cmd.InitDepot(c.String("depot-path"))
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}

}
