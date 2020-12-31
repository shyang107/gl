package main

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/shyang107/paw"
	"github.com/urfave/cli"
)

func getPath(c *cli.Context) string {
	path := c.Args().Get(0)
	if len(path) == 0 {
		path = "."
	}
	if paw.HasPrefix(path, "~") {
		path, err = homedir.Expand(path)
	} else {
		path, err = filepath.Abs(path)
	}
	if err != nil || !paw.IsExist(path) {
		// paw.Error.Printf("%q error: %v", path, err)
		paw.Error.Printf("%q does not exist or error: %v", path, err)
		os.Exit(1)
	}
	return path
}
