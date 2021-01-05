package main

import "github.com/shyang107/paw/filetree"

func ckView(opt *gloption, pdopt *filetree.PrintDirOption) {

	if opt.isListTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.OutOpt = filetree.PListTreeView
	} else if opt.isTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.OutOpt = filetree.PTreeView
	} else if opt.isTable {
		pdopt.OutOpt = filetree.PTableView
	} else if opt.isLevel {
		pdopt.OutOpt = filetree.PLevelView
	} else if opt.isClassify {
		pdopt.OutOpt = filetree.PClassifyView
	}

	if opt.isExtended {
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
}
