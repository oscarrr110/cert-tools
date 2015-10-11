package cmd

import (
	"fmt"
	"os"
	"github.com/oscarrr110/cert-tools/depot"
)

var (
	d *depot.FileDepot
)

func InitDepot(path string) error {
	if d == nil {
		var err error
		fmt.Println("init depot")
		if d, err = depot.NewFileDepot(path); err != nil {
			fmt.Println("init depot error")
			return err
		}

		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}
