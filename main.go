package main

import (
	"fmt"
	"os"

	"github.com/spf13/cast"

	"github.com/shyang107/paw"
	_ "github.com/shyang107/paw"
	"github.com/urfave/cli"
)

const (
	version = "0.0.7.2"
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
	app.UsageText = "web-gl command [command options] [arguments...]"
	app.Version = version
	// app.Compiled = time.Now()
	app.Compiled = cast.ToTime("2021-02-1")
	app.Authors = []*cli.Author{
		{
			Name:  "Shuhhua Yang",
			Email: "shyang107@gmail.com",
		},
	}
	app.ArgsUsage = "[path]"

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s @ %v\n", c.App.Name, paw.NewEXAColor("sb").Sprint("gl"+c.App.Version), paw.NewEXAColor("da").Sprint(c.App.Compiled.Format("Jan 2, 2006")))
	}

	// app.EnableBashCompletion = true

	app.UseShortOptionHandling = true

	app.Commands = []*cli.Command{
		{
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
		&isFieldINodeFlag, &isFieldLinksFlag,
		// &isFieldPermissionsFlag,
		// &isFieldSizeFlag,
		&isFieldBlocksFlag,
		// &isFieldUserFlag, &isFieldGroupFlag,
		&isModifiedFlag, &isAccessedFlag, &isCreatedFlag,
		&isFieldGitFlag,
		&isExtendedFlag,
		&isNoSortFlag, &isReverseFlag, &sortByFieldFlag, &isSortByNameFlag, &isSortBySizeFlag, &isSortByMTimeFlag,
		&isGroupedFlag,
	}

	app.Action = appAction
}

func main() {
	// start := time.Now()

	err := app.Run(os.Args)
	if err != nil {
		paw.Error.Printf("run '%s' failed, error:%v", app.Name, err)
		os.Exit(1)
	}

	// elapsedTime := time.Since(start)
	// fmt.Println()
	// fmt.Println("Total time for excution:", elapsedTime.String())
}
