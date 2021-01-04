package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/shyang107/paw/filetree"

	"github.com/shyang107/paw"
	_ "github.com/shyang107/paw"
	"github.com/urfave/cli"
)

const (
	version = "0.0.4"
)

var (
	app = cli.NewApp()
)

func init() {

	paw.GologInit(os.Stdout, os.Stdout, os.Stderr, false)

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v", "V"},
		Usage:   "print only the version",
	}

	app.Name = "gl"
	app.Usage = "list directory (excluding hidden items) in color view."
	app.Version = version
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Shuhhua Yang",
			Email: "shyang107@gmail.com",
		},
	}
	app.ArgsUsage = "[path]"

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s @ %v\n", c.App.Name, color.New(color.FgHiGreen).Sprint(c.App.Version), filetree.NewEXAColor("da").Sprint(c.App.Compiled.Format("Jan 2, 2006")))
	}

	app.Commands = []*cli.Command{
		&cli.Command{
			Name:    "version",
			Aliases: []string{"v", "V"},
			Usage:   "print only the version",
			Action: func(c *cli.Context) error {
				cli.ShowVersion(c)
				return nil
			},
		},
	}

	app.Flags = []cli.Flag{
		&listFlag, &listTreeFlag, &treeFlag, &tableFlag, &levelFlag, &clsassifyFlag, &depthFlag, &recurseFlag,
		&allFilesFlag, &includePatternFlag, &excludePatternFlag,
		&isNoEmptyDirsFlag, &isJustDirsFlag, &isJustFilesFlag,
		&isNoSortFlag, &isReverseFlag, &isSortByNameFlag, &isSortBySizeFlag, &isSortByMTimeFlag,
		&isGroupedFlag,
		&isExtendedFlag,
	}

	app.Action = appAction
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		paw.Error.Printf("run '%s' failed, error:%v", app.Name, err)
		os.Exit(1)
	}
}
