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
	version = "0.0.2"
)

var (
	app = cli.NewApp()

	path           string
	isList         bool
	isListTree     bool
	isTree         bool
	isTable        bool
	isLevel        bool
	isClassify     bool
	isRecurse      bool
	depth          int
	isAllFiles     bool
	includePattern string
	excludePattern string

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
	clsassifyFlag = cli.BoolFlag{
		Name:        "classify",
		Aliases:     []string{"F"},
		Value:       false,
		Usage:       "display type indicator by file names",
		Destination: &isClassify,
	}
	depthFlag = cli.IntFlag{
		Name:        "depth",
		Aliases:     []string{"d"},
		Value:       0,
		Usage:       "print out in the level view",
		Destination: &depth,
	}
	recurseFlag = cli.BoolFlag{
		Name:        "recurse",
		Aliases:     []string{"R"},
		Value:       false,
		Usage:       "recurse into directories (equivalent to --depth=-1)",
		Destination: &isRecurse,
	}
	allFilesFlag = cli.BoolFlag{
		Name:        "all",
		Aliases:     []string{"a"},
		Value:       false,
		Usage:       "show all file including hidden files",
		Destination: &isAllFiles,
	}
	includePatternFlag = cli.StringFlag{
		Name:        "include",
		Aliases:     []string{"n"},
		Value:       "",
		Usage:       "set regex `pattern` to include some files, applied to file only",
		Destination: &includePattern,
	}
	excludePatternFlag = cli.StringFlag{
		Name:        "exclude",
		Aliases:     []string{"x"},
		Value:       "",
		Usage:       "set regex `pattern` to exclude some files, applied to file only",
		Destination: &excludePattern,
	}

	pdopt = filetree.NewPrintDirOption()

	err error
)

type patflag int

const (
	allFlag patflag = 1 << iota
	includeFlag
	excludeFlag
	allincludeFlag      = allFlag | includeFlag
	allexcludeFlag      = allFlag | excludeFlag
	allinAndexcludeFlag = allFlag | includeFlag | excludeFlag
	inAndexcludeFlag    = includeFlag | excludeFlag
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
		&listFlag, &listTreeFlag, &treeFlag, &tableFlag, &levelFlag, &clsassifyFlag, &depthFlag, &recurseFlag, &allFilesFlag, &includePatternFlag, &excludePatternFlag,
	}

	app.Action = func(c *cli.Context) error {
		path = getPath(c)

		// isList,	isListTree, isTree, isTable, isLevel, depth
		ckView()

		// pattern
		pflag := getpatflag()
		switch pflag {
		case allFlag:
			optAllFiles()
		case includeFlag:
			optInclude()
		case excludeFlag:
			optExclude()
		case allincludeFlag:
			optAllInclude()
		case allexcludeFlag:
			optAllExclude()
		case allinAndexcludeFlag:
			optAllInAndExclude()
		case inAndexcludeFlag:
			optInAndExclude()
		}

		err := filetree.PrintDir(os.Stdout, path, pdopt, "")
		if err != nil {
			paw.Error.Printf("get file list from %q failed, error: %v", path, err)
			os.Exit(1)
		}

		return nil
	}
}

func getpatflag() (pflag patflag) {
	if isAllFiles && len(excludePattern) == 0 && len(includePattern) == 0 {
		pflag = allFlag
		goto END
	}
	if isAllFiles && len(excludePattern) > 0 && len(includePattern) == 0 {
		pflag = allexcludeFlag
		goto END
	}
	if isAllFiles && len(excludePattern) > 0 && len(includePattern) > 0 {
		pflag = allinAndexcludeFlag
		goto END
	}
	if isAllFiles && len(excludePattern) == 0 && len(includePattern) > 0 {
		pflag = allincludeFlag
		goto END
	}

	if !isAllFiles && len(excludePattern) > 0 && len(includePattern) == 0 {
		pflag = excludeFlag
		goto END
	}
	if !isAllFiles && len(excludePattern) > 0 && len(includePattern) > 0 {
		pflag = inAndexcludeFlag
		goto END
	}
	if !isAllFiles && len(excludePattern) == 0 && len(includePattern) > 0 {
		pflag = includeFlag
		goto END
	}
END:
	return pflag
}

func optAllInAndExclude() {
	ren, err := regexp.Compile(includePattern)
	if err != nil {
		paw.Error.Printf("including pattern: %q, error: %v", ren.String(), err)
		os.Exit(1)
	}
	rex, err := regexp.Compile(excludePattern)
	if err != nil {
		paw.Error.Printf("excluding pattern: %q, error: %v", rex.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			if !ren.MatchString(f.BaseName) && rex.MatchString(f.BaseName) {
				return filetree.SkipThis
			}
		}
		return nil
	}
}
func optInAndExclude() {
	ren, err := regexp.Compile(includePattern)
	if err != nil {
		paw.Error.Printf("including pattern: %q, error: %v", ren.String(), err)
		os.Exit(1)
	}
	rex, err := regexp.Compile(excludePattern)
	if err != nil {
		paw.Error.Printf("excluding pattern: %q, error: %v", rex.String(), err)
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
		if !f.IsDir() {
			if !ren.MatchString(f.BaseName) && rex.MatchString(f.BaseName) {
				return filetree.SkipThis
			}
		}
		return nil
	}
}
func optAllInclude() {
	re, err := regexp.Compile(includePattern)
	if err != nil {
		paw.Error.Printf("including pattern: %q, error: %v", re.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			if !re.MatchString(f.BaseName) {
				return filetree.SkipThis
			}
		}
		return nil
	}
}
func optInclude() {
	re, err := regexp.Compile(includePattern)
	if err != nil {
		paw.Error.Printf("including pattern: %q, error: %v", re.String(), err)
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
		if !f.IsDir() {
			if !re.MatchString(f.BaseName) {
				return filetree.SkipThis
			}
		}
		return nil
	}
}
func optAllExclude() {
	re, err := regexp.Compile(excludePattern)
	if err != nil {
		paw.Error.Printf("excluding pattern: %q, error: %v", re.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			if re.MatchString(f.BaseName) {
				return filetree.SkipThis
			}
		}
		return nil
	}
}
func optExclude() {
	re, err := regexp.Compile(excludePattern)
	if err != nil {
		paw.Error.Printf("excluding pattern: %q, error: %v", re.String(), err)
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
		if !f.IsDir() {
			if re.MatchString(f.BaseName) {
				return filetree.SkipThis
			}
		}
		return nil
	}
}
func optAllFiles() {
	pdopt.Ignore = func(f *filetree.File, e error) error {
		return nil
	}
}

func ckView() {
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
	} else if isClassify {
		pdopt.OutOpt = filetree.PClassifyView
	}
	// else {
	// 	pdopt.OutOpt = filetree.PListView
	// }

	if isRecurse {
		depth = -1
	}

	pdopt.Depth = depth
}

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

func main() {
	err := app.Run(os.Args)
	if err != nil {
		paw.Error.Printf("run '%s' failed, error:%v", app.Name, err)
		os.Exit(1)
	}
}
