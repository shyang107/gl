package main

import (
	"os"

	"github.com/shyang107/paw"
	"github.com/shyang107/paw/filetree"
	"github.com/urfave/cli"
)

var appAction = func(c *cli.Context) error {
	opt.path = getPath(c)

	// isList,	isListTree, isTree, isTable, isLevel, depth
	ckView(opt, pdopt)

	// pattern
	pflag := getpatflag(opt)
	switch pflag {
	case allFlag:
		optAllFiles(opt, pdopt)
	case includeFlag:
		optInclude(opt, pdopt)
	case excludeFlag:
		optExclude(opt, pdopt)
	case allincludeFlag:
		optAllInclude(opt, pdopt)
	case allexcludeFlag:
		optAllExclude(opt, pdopt)
	case allinAndexcludeFlag:
		optAllInAndExclude(opt, pdopt)
	case inAndexcludeFlag:
		optInAndExclude(opt, pdopt)
	}

	sortOpt := getSortOption(opt)

	err := filetree.PrintDir(os.Stdout, opt.path, opt.isGrouped, pdopt, sortOpt, "")
	if err != nil {
		paw.Error.Printf("get file list from %q failed, error: %v", opt.path, err)
		os.Exit(1)
	}

	return nil
}
