package main

import (
	"github.com/shyang107/paw/filetree"
)

var (
	byMTime filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.ModifiedTime().Before(fj.ModifiedTime())
	}
	bySize filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.Size < fj.Size
	}
)

func getSortOption(opt *gloption) *filetree.PrintDirSortOption {
	if opt.isNoSort {
		return &filetree.PrintDirSortOption{
			IsSort: !opt.isNoSort,
		}
	}

	// var sortOpt *filetree.PrintDirSortOption
	if opt.isSortBySize {
		// paw.Info.Println("opt.isSortBySize", opt.isSortBySize)
		// paw.Info.Println("  opt.isReverse", opt.isReverse)
		if opt.isReverse {
			return &filetree.PrintDirSortOption{
				IsSort:  true,
				SortWay: filetree.PDSortByReverseSize,
			}
		} else {
			return &filetree.PrintDirSortOption{
				IsSort:  true,
				SortWay: filetree.PDSortBySize,
			}
		}
	}
	if opt.isSortByMTime {
		// paw.Info.Println("opt.isSortByMTime", opt.isSortByMTime)
		// paw.Info.Println("  opt.isReverse", opt.isReverse)
		if opt.isReverse {
			return &filetree.PrintDirSortOption{
				IsSort:  true,
				SortWay: filetree.PDSortByReverseMtime,
			}
		} else {
			return &filetree.PrintDirSortOption{
				IsSort:  true,
				SortWay: filetree.PDSortByMtime,
			}
		}
	}
	// if opt.isSortByName { //default
	opt.isSortByName = true
	// paw.Info.Println("opt.isSortByName", opt.isSortByName)
	// paw.Info.Println("  opt.isReverse", opt.isReverse)
	if opt.isReverse {
		return &filetree.PrintDirSortOption{
			IsSort:  true,
			SortWay: filetree.PDSortByReverseName,
		}
	} else {
		return &filetree.PrintDirSortOption{
			IsSort:  true,
			SortWay: filetree.PDSortByName,
		}
	}
	// }

	return nil
}
