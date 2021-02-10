package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cast"

	"github.com/shyang107/paw"
	_ "github.com/shyang107/paw"
	"github.com/urfave/cli"
)

const (
	version = "0.0.7.4"
)

var (
	app         = cli.NewApp()
	programName string
	lg          = paw.Logger
)

func init() {
	programName, err := os.Executable()
	if err != nil {
		programName = os.Args[0]
	}
	programName = filepath.Base(programName)

	paw.GologInit(os.Stdout, os.Stderr, os.Stderr, false)

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app.Name = "gl"
	app.Usage = "list directory (excluding hidden items) in color view."
	app.UsageText = "web-gl command [command options] [arguments...]"
	app.Version = version
	// app.Compiled = time.Now()
	app.Compiled = cast.ToTime("2021-02-10")
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
			Aliases: []string{"v"},
			Usage:   "print only the version",
			Action: func(c *cli.Context) error {
				cli.ShowVersion(c)
				return nil
			},
		},
	}

	app.Flags = []cli.Flag{
		&verboseFlag,
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
		fatal("run '%s' failed, error:%v", app.Name, err)
	}

	// elapsedTime := time.Since(start)
	// fmt.Println()
	// fmt.Println("Total time for excution:", elapsedTime.String())
}

func info(f string, args ...interface{}) {
	if opt.isVerbose {
		paw.Info.Printf(programName + ": " + fmt.Sprintf(f, args...) + "\n")
	}
	// fmt.Fprintf(os.Stderr, programName+": "+fmt.Sprintf(f, args...)+"\n")
}

func stderr(f string, args ...interface{}) {
	paw.Error.Printf(programName + ": " + fmt.Sprintf(f, args...) + "\n")
	// fmt.Fprintf(os.Stderr, programName+": "+fmt.Sprintf(f, args...)+"\n")
}

func fatal(f string, args ...interface{}) {
	stderr(f, args...)
	os.Exit(1)
}

func warning(f string, args ...interface{}) {
	if opt.isVerbose {
		paw.Warning.Printf(programName + ": " + fmt.Sprintf(f, args...) + "\n")
		// stderr(f, args...)
	}
}
