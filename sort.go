package main

import (
	"strings"

	"github.com/shyang107/paw"
	"github.com/shyang107/paw/filetree"
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

func getSortOption(opt *gloption) *filetree.PDSortOption {
	if opt.isNoSort {
		return &filetree.PDSortOption{
			IsSort: !opt.isNoSort,
		}
	}
	var sflag = strings.ToLower(opt.sortByField)
	if len(sflag) > 0 {
		if strings.HasSuffix(sflag, "r") {
			opt.isReverse = true
		}
		if opt.isReverse && !strings.HasSuffix(sflag, "r") {
			sflag += "r"
		}
		if flag, ok := sortMapFlag[sflag]; ok {
			return &filetree.PDSortOption{
				IsSort:  true,
				SortWay: flag,
			}
		} else {
			paw.Error.Printf("%q is not allowed; so, sort by name in increasing order!\n", opt.sortByField)
			flag = sortMapFlag["name"]
			return &filetree.PDSortOption{
				IsSort:  true,
				SortWay: flag,
			}
		}
	}

	// var sortOpt *filetree.PDSortOption
	if opt.isSortBySize {
		sflag = "size"
	} else if opt.isSortByMTime {
		sflag = "mtime"
	} else {
		opt.isSortByName = true
		sflag = "name"
	}
	if opt.isReverse {
		sflag += "r"
	}

	return &filetree.PDSortOption{
		IsSort:  true,
		SortWay: sortMapFlag[sflag],
	}
}
