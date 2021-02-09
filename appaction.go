package main

import (
	"os"

	"github.com/shyang107/paw/filetree"
	"github.com/urfave/cli"
)

var appAction = func(c *cli.Context) error {

	checkArgs(c, pdopt)

	// isList,	isListTree, isTree, isTable, isLevel, depth
	checkView(opt, pdopt)

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

	// sortOpt := getSortOption(opt)

	err, _ := filetree.PrintDir(os.Stdout, opt.path, pdopt, "")
	if err != nil {
		fatal("get file list from %q failed, error: %v", opt.path, err)
	}

	return nil
}
