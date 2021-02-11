package main

import "github.com/shyang107/paw/filetree"

func checkView(opt *gloption, pdopt *filetree.PrintDirOption) {

	if opt.isListTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.ViewFlag = filetree.PListTreeView
		lg.WithField("view", "ListTree").Trace()
	} else if opt.isTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.ViewFlag = filetree.PTreeView
		lg.WithField("view", "Tree").Trace()
	} else if opt.isTable {
		pdopt.ViewFlag = filetree.PTableView
		lg.WithField("view", "Table").Trace()
	} else if opt.isLevel {
		pdopt.ViewFlag = filetree.PLevelView
		lg.WithField("view", "Level").Trace()
	} else if opt.isClassify {
		pdopt.ViewFlag = filetree.PClassifyView
		lg.WithField("view", "Clssify").Trace()
	} else if opt.isList {
		lg.WithField("view", "List").Trace()
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
	lg.WithField("depth", pdopt.Depth).Trace()
}
