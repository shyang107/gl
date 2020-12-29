package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/shyang107/paw/filetree"

	"github.com/shyang107/paw"
	_ "github.com/shyang107/paw"
	"github.com/urfave/cli"
)

const (
	version = "0.0.1"
)

var (
	app = cli.NewApp()

	path          string
	isList        bool
	isListTree    bool
	isTree        bool
	isTable       bool
	isLevel       bool
	depth         int
	isAllFiles    bool
	ignorePattern string

	listFlag = cli.BoolFlag{
		Name:        "list",
		Aliases:     []string{"l"},
		Value:       true,
		Usage:       "print out in list view",
		Destination: &isList,
	}
	listTreeFlag = cli.BoolFlag{
		Name:        "listtree",
		Aliases:     []string{"t"},
		Value:       false,
		Usage:       "print out in the view of combining list and tree",
		Destination: &isListTree,
	}
	treeFlag = cli.BoolFlag{
		Name:        "tree",
		Aliases:     []string{"T"},
		Value:       false,
		Usage:       "print out in the tree view",
		Destination: &isTree,
	}
	tableFlag = cli.BoolFlag{
		Name:        "table",
		Aliases:     []string{"b"},
		Value:       false,
		Usage:       "print out in the table view",
		Destination: &isTable,
	}
	levelFlag = cli.BoolFlag{
		Name:        "level",
		Aliases:     []string{"L"},
		Value:       false,
		Usage:       "print out in the level view",
		Destination: &isLevel,
	}
	depthFlag = cli.IntFlag{
		Name:        "depth",
		Aliases:     []string{"d"},
		Value:       0,
		Usage:       "print out in the level view",
		Destination: &depth,
	}
	allFilesFlag = cli.BoolFlag{
		Name:        "all",
		Aliases:     []string{"a"},
		Value:       false,
		Usage:       "show all file including hidden files",
		Destination: &isAllFiles,
	}
	ignorePatternFlag = cli.StringFlag{
		Name:        "ignore-pattern",
		Aliases:     []string{"p"},
		Value:       "",
		Usage:       "set pattern of `regexp` to ignore some files",
		Destination: &ignorePattern,
	}

	pdopt = filetree.NewPrintDirOption()

	err error
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
	app.ArgsUsage = "[directory]"

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
		&listFlag, &listTreeFlag, &treeFlag, &tableFlag, &levelFlag, &depthFlag, &allFilesFlag, &ignorePatternFlag,
	}

	app.Action = func(c *cli.Context) error {
		path = c.Args().Get(0)
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

		// if isList {
		// 	pdopt.OutOpt = filetree.PListView
		// } else
		if isListTree {
			if depth == 0 {
				depth = -1
			}
			pdopt.OutOpt = filetree.PListTreeView
		} else if isTree {
			if depth == 0 {
				depth = -1
			}
			pdopt.OutOpt = filetree.PTreeView
		} else if isTable {
			pdopt.OutOpt = filetree.PTableView
		} else if isLevel {
			pdopt.OutOpt = filetree.PLevelView
		}
		// else {
		// 	pdopt.OutOpt = filetree.PListView
		// }

		if isAllFiles {
			pdopt.Ignore = func(f *filetree.File, e error) error {
				return nil
			}
		}
		if len(ignorePattern) > 0 {
			re, err := regexp.Compile(ignorePattern)
			if err != nil {
				paw.Error.Printf("%s, error: %v", re.String(), err)
				os.Exit(1)
			}
			pdopt.Ignore = func(f *filetree.File, e error) error {
				if err != nil {
					return err
				}
				_, file := filepath.Split(f.Path)
				if paw.HasPrefix(file, ".") {
					return filetree.SkipThis
				}
				if re.MatchString(f.BaseName) {
					return filetree.SkipThis
				}
				return nil
			}
		}

		if isAllFiles && len(ignorePattern) > 0 {
			re, err := regexp.Compile(ignorePattern)
			if err != nil {
				paw.Error.Printf("%s, error: %v", re.String(), err)
				os.Exit(1)
			}
			pdopt.Ignore = func(f *filetree.File, e error) error {
				if err != nil {
					return err
				}
				if re.MatchString(f.BaseName) {
					return filetree.SkipThis
				}
				return nil
			}
		}

		pdopt.Depth = depth

		err := filetree.PrintDir(os.Stdout, path, pdopt, "")
		if err != nil {
			paw.Error.Printf("get file list from %q failed, error:%v", path, err)
			os.Exit(1)
		}

		return nil
	}
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		paw.Error.Printf("run '%s' failed, error:%v", app.Name, err)
		os.Exit(1)
	}
}
