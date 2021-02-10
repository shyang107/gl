package main

import "github.com/shyang107/paw/filetree"

func checkView(opt *gloption, pdopt *filetree.PrintDirOption) {

	if opt.isListTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.OutOpt = filetree.PListTreeView
		lg.WithField("view", "ListTree").Trace()
	} else if opt.isTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.OutOpt = filetree.PTreeView
		lg.WithField("view", "Tree").Trace()
	} else if opt.isTable {
		pdopt.OutOpt = filetree.PTableView
		lg.WithField("view", "Table").Trace()
	} else if opt.isLevel {
		pdopt.OutOpt = filetree.PLevelView
		lg.WithField("view", "Level").Trace()
	} else if opt.isClassify {
		pdopt.OutOpt = filetree.PClassifyView
		lg.WithField("view", "Clssify").Trace()
	} else if opt.isList {
		lg.WithField("view", "List").Trace()
	}

	if opt.isExtended {
		lg.Trace("show extended attributes")
		switch {
		case pdopt.OutOpt == filetree.PLevelView:
			pdopt.OutOpt = filetree.PLevelExtendView
		case pdopt.OutOpt == filetree.PTableView:
			pdopt.OutOpt = filetree.PTableExtendView
		case pdopt.OutOpt == filetree.PTreeView:
			pdopt.OutOpt = filetree.PTreeExtendView
		case pdopt.OutOpt == filetree.PListTreeView:
			pdopt.OutOpt = filetree.PListTreeExtendView
		default:
			pdopt.OutOpt = filetree.PListExtendView
		}
	}

	if opt.isRecurse {
		opt.depth = -1
	}

	pdopt.Depth = opt.depth
	lg.WithField("depth", pdopt.Depth).Trace()
}
