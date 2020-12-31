package main

import (
	"github.com/shyang107/paw/filetree"
	"github.com/urfave/cli"
)

type gloption struct {
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
}

var (
	opt   = new(gloption)
	pdopt = filetree.NewPrintDirOption()
	err   error

	listFlag = cli.BoolFlag{
		Name:        "list",
		Aliases:     []string{"l"},
		Value:       true,
		Usage:       "print out in list view",
		Destination: &opt.isList,
	}
	listTreeFlag = cli.BoolFlag{
		Name:        "listtree",
		Aliases:     []string{"t"},
		Value:       false,
		Usage:       "print out in the view of combining list and tree",
		Destination: &opt.isListTree,
	}
	treeFlag = cli.BoolFlag{
		Name:        "tree",
		Aliases:     []string{"T"},
		Value:       false,
		Usage:       "print out in the tree view",
		Destination: &opt.isTree,
	}
	tableFlag = cli.BoolFlag{
		Name:        "table",
		Aliases:     []string{"b"},
		Value:       false,
		Usage:       "print out in the table view",
		Destination: &opt.isTable,
	}
	levelFlag = cli.BoolFlag{
		Name:        "level",
		Aliases:     []string{"L"},
		Value:       false,
		Usage:       "print out in the level view",
		Destination: &opt.isLevel,
	}
	clsassifyFlag = cli.BoolFlag{
		Name:        "classify",
		Aliases:     []string{"F"},
		Value:       false,
		Usage:       "display type indicator by file names",
		Destination: &opt.isClassify,
	}
	depthFlag = cli.IntFlag{
		Name:        "depth",
		Aliases:     []string{"d"},
		Value:       0,
		Usage:       "print out in the level view",
		Destination: &opt.depth,
	}
	recurseFlag = cli.BoolFlag{
		Name:        "recurse",
		Aliases:     []string{"R"},
		Value:       false,
		Usage:       "recurse into directories (equivalent to --depth=-1)",
		Destination: &opt.isRecurse,
	}
	allFilesFlag = cli.BoolFlag{
		Name:        "all",
		Aliases:     []string{"a"},
		Value:       false,
		Usage:       "show all file including hidden files",
		Destination: &opt.isAllFiles,
	}
	includePatternFlag = cli.StringFlag{
		Name:        "include",
		Aliases:     []string{"n"},
		Value:       "",
		Usage:       "set regex `pattern` to include some files, applied to file only",
		Destination: &opt.includePattern,
	}
	excludePatternFlag = cli.StringFlag{
		Name:        "exclude",
		Aliases:     []string{"x"},
		Value:       "",
		Usage:       "set regex `pattern` to exclude some files, applied to file only",
		Destination: &opt.excludePattern,
	}
)
