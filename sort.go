package main

import (
	"strings"

	"github.com/shyang107/paw/filetree"
	"github.com/sirupsen/logrus"
)

var (
	sortMapFlag = map[string]filetree.PDSortFlag{
		"inode":     filetree.PDSortByINode,
		"links":     filetree.PDSortByLinks,
		"size":      filetree.PDSortBySize,
		"blocks":    filetree.PDSortByBlocks,
		"modified":  filetree.PDSortByMTime,
		"mtime":     filetree.PDSortByMTime,
		"accessed":  filetree.PDSortByATime,
		"atime":     filetree.PDSortByATime,
		"created":   filetree.PDSortByCTime,
		"ctime":     filetree.PDSortByCTime,
		"name":      filetree.PDSortByName,
		"inoder":    filetree.PDSortByINodeR,
		"linksr":    filetree.PDSortByLinksR,
		"sizer":     filetree.PDSortBySizeR,
		"blocksr":   filetree.PDSortByBlocksR,
		"modifiedr": filetree.PDSortByMTimeR,
		"mtimer":    filetree.PDSortByMTimeR,
		"accessedr": filetree.PDSortByATimeR,
		"atimer":    filetree.PDSortByATimeR,
		"createdr":  filetree.PDSortByCTimeR,
		"ctimer":    filetree.PDSortByCTimeR,
		"namer":     filetree.PDSortByNameR,
	}
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
		// paw.Logger.WithField("sflag", sflag).Info()
		sopt.Reverse = opt.isReverse
		if flag, ok := sortMapFlag[sflag]; ok {
			sopt.SortWay = flag
		} else {
			stderr("%q is not allowed; so, sort by name in increasing order!", opt.sortByField)
			sflag = "name"
			flag = sortMapFlag[sflag]
			sopt.SortWay = flag
		}
	} else {
		sflag = "name"
		sopt.SortWay = sortMapFlag[sflag]
	}
	lg.WithFields(logrus.Fields{
		"isSort":    sopt.IsSort,
		"isReverse": sopt.Reverse,
		"sortBy":    sflag,
	}).Trace()
	return sopt
}
