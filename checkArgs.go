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
		path, err := filepath.Abs(".")
		if err != nil {
			paw.Error.Println(err)
		}
		pdopt.SetRoot(path)
		info("[checkArgs] no argument, suppose root is %q", path)
	case 1:
		path, err := filepath.Abs(c.Args().Get(0))
		if err != nil {
			paw.Error.Println(err)
		}
		info("[checkArgs] one argument, root is %q", path)
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
	// 	// paw.Logger.WithField("path", path).Info("One")
	default: // > 1
		info("[checkArgs] multi-arguments")
		for i := 0; i < c.NArg(); i++ {
			// paw.Logger.WithField("args", c.Args().Get(i)).Info()
			path, err := filepath.Abs(c.Args().Get(i))
			if err != nil {
				paw.Error.Println(err)
				continue
			}
			info("[checkArgs] multi-arguments, argument %d is %q", i+1, path)
			pdopt.AddPath(path)
			// paw.Logger.WithField("path", path).Info("Multi")
		}
	}
}
