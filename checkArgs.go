package main

import (
	"os"
	"path/filepath"

	"github.com/shyang107/paw"
	"github.com/shyang107/paw/filetree"
	"github.com/urfave/cli"
)

func checkArgs(c *cli.Context, pdopt *filetree.PrintDirOption) {

	switch c.NArg() {
	case 0:
		lg.WithField("arg", c.Args().Get(0)).Trace("no argument")
		path, err := filepath.Abs(".")
		if err != nil {
			paw.Error.Println(err)
		}
		pdopt.SetRoot(path)
	case 1:
		lg.WithField("arg", c.Args().Get(0)).Trace("no argument")
		path, err := filepath.Abs(c.Args().Get(0))
		if err != nil {
			paw.Error.Println(err)
		}
		fi, err := os.Stat(path)
		if err != nil {
			paw.Error.Println(err)
			os.Exit(1)
		}
		if fi.IsDir() {
			pdopt.SetRoot(path)
		} else {
			pdopt.AddPath(path)
		}
	default: // > 1
		lg.WithField("arg", c.Args()).Trace("multi-arguments")
		for i := 0; i < c.NArg(); i++ {
			// paw.Logger.WithField("args", c.Args().Get(i)).Info()
			path, err := filepath.Abs(c.Args().Get(i))
			if err != nil {
				paw.Error.Println(err)
				continue
			}
			pdopt.AddPath(path)
		}
	}
}
