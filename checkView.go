package main

import "github.com/shyang107/paw/filetree"

func checkView(opt *gloption, pdopt *filetree.PrintDirOption) {

	if opt.isListTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.OutOpt = filetree.PListTreeView
		info("[checkView] pdopt.OutOpt: ListTree view")
	} else if opt.isTree {
		if opt.depth == 0 {
			opt.depth = -1
		}
		pdopt.OutOpt = filetree.PTreeView
		info("[checkView] pdopt.OutOpt: Tree view")
	} else if opt.isTable {
		pdopt.OutOpt = filetree.PTableView
	} else if opt.isLevel {
		pdopt.OutOpt = filetree.PLevelView
		info("[checkView] pdopt.OutOpt: Table view")
	} else if opt.isClassify {
		pdopt.OutOpt = filetree.PClassifyView
		info("[checkView] pdopt.OutOpt: Clssify view")
	} else if opt.isList {
		info("[checkView] pdopt.OutOpt: List view")
	}

	if opt.isExtended {
		info("[checkView] show extended attributes")
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
	info("[checkView] pdopt.Depth is %d", pdopt.Depth)
}
