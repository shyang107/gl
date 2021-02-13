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
	} else {
		lg.SetLevel(logrus.WarnLevel)
	}

	checkArgs(c, pdopt)

	// isList,	isListTree, isTree, isTable, isLevel, depth
	checkView(opt, pdopt)

	// pattern
	pdopt.Ignore = getPatternflag(opt).Ignore(opt)

	pdopt.FieldFlag = getFieldFlag(opt)
	pdopt.SortOpt = getSortOption(opt)
	pdopt.FiltOpt = getFiltOption(opt)

	err, _ := filetree.PrintDir(os.Stdout, opt.path, opt.isGrouped, pdopt, "")
	if err != nil {
		fatal("get file list from %q failed, error: %v", opt.path, err)
	}

	return nil
}
