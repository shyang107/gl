package main

import (
	"github.com/shyang107/paw/filetree"
	"github.com/sirupsen/logrus"
)

func checkView(opt *gloption, pdopt *filetree.PrintDirOption) {

	if opt.isListTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.ViewFlag = filetree.PListTreeView
	} else if opt.isTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.ViewFlag = filetree.PTreeView
	} else if opt.isTable {
		pdopt.ViewFlag = filetree.PTableView
	} else if opt.isLevel {
		pdopt.ViewFlag = filetree.PLevelView
	} else if opt.isClassify {
		pdopt.ViewFlag = filetree.PClassifyView
	} else if opt.isList {
		pdopt.ViewFlag = filetree.PListView
	}

	if opt.isExtended {
		lg.Trace("show extended attributes")
		var view filetree.PDViewFlag
		switch pdopt.ViewFlag {
		case filetree.PLevelView:
			view = filetree.PLevelExtendView
		case filetree.PTableView:
			view = filetree.PTableExtendView
		case filetree.PTreeView:
			view = filetree.PTreeExtendView
		case filetree.PListTreeView:
			view = filetree.PListTreeExtendView
		default:
			view = filetree.PListExtendView
		}
		pdopt.ViewFlag = view
	}
	if opt.isRecurse {
		opt.depth = -1
	}

	pdopt.Depth = opt.depth

	lg.WithFields(logrus.Fields{
		"view":  pdopt.ViewFlag.String(),
		"depth": pdopt.Depth,
	}).Trace("view")
}
