package main

import (
	"strings"

	"github.com/shyang107/paw/filetree"
	"github.com/sirupsen/logrus"
)

func getSortOption(opt *gloption) (sopt *filetree.PDSortOption) {
	sopt = &filetree.PDSortOption{
		IsSort:  true,
		Reverse: false,
		// SortWay: filetree.PDSortByName, // PDSortFlag
	}

	if opt.isNoSort {
		sopt.IsSort = false
		return sopt
	}

	var sflag = strings.ToLower(opt.sortByField)
	if opt.isSortBySize {
		sflag = "size"
	} else if opt.isSortByMTime {
		sflag = "mtime"
	} else if opt.isSortByName {
		sflag = "name"
		opt.isSortByName = true
	}

	if len(sflag) > 0 {
		if strings.HasSuffix(sflag, "r") {
			opt.isReverse = true
		}
		if opt.isReverse && !strings.HasSuffix(sflag, "r") {
			sflag += "r"
		}
		sopt.Reverse = opt.isReverse
	}
	sopt.SortFlag = sopt.SortFlag.GetFlag(sflag)

	lg.WithFields(logrus.Fields{
		"isSort":    sopt.IsSort,
		"isReverse": sopt.Reverse,
		"sort":      sopt.SortFlag.String(),
	}).Trace("sort")
	return sopt
}
