package main

import (
	"os"

	"github.com/shyang107/paw/filetree"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var appAction = func(c *cli.Context) error {

	if opt.isVerbose {
		pdopt.EnableTrace(opt.isVerbose)
		lg.SetLevel(logrus.TraceLevel)
	}

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

	pdopt.FieldFlag = getFieldFlag(opt)
	pdopt.SortOpt = getSortOption(opt)
	pdopt.FiltOpt = getFiltOption(opt)

	err, _ := filetree.PrintDir(os.Stdout, opt.path, opt.isGrouped, pdopt, "")
	if err != nil {
		fatal("get file list from %q failed, error: %v", opt.path, err)
	}

	return nil
}
