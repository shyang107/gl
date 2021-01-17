package main

import (
	"github.com/shyang107/paw"
	"github.com/shyang107/paw/filetree"
)

var (
	bySize filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.Size < fj.Size
	}
	bySizeR filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.Size > fj.Size
	}

	byMTime filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.ModifiedTime().Before(fj.ModifiedTime())
	}
	byMTimeR filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.ModifiedTime().After(fj.ModifiedTime())
	}
	byATime filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.AccessedTime().Before(fj.AccessedTime())
	}
	byATimeR filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.AccessedTime().After(fj.AccessedTime())
	}
	byCTime filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.CreatedTime().Before(fj.CreatedTime())
	}
	byCTimeR filetree.FilesBy = func(fi, fj *filetree.File) bool {
		return fi.CreatedTime().After(fj.CreatedTime())
	}

	byName filetree.FilesBy = func(fi, fj *filetree.File) bool {
		// if fi.IsDir() && fj.IsFile() {
		// 	return true
		// } else if fi.IsFile() && fj.IsDir() {
		// 	return false
		// }
		return paw.ToLower(fi.BaseName) < paw.ToLower(fj.BaseName)
	}
	byNameR filetree.FilesBy = func(fi, fj *filetree.File) bool {
		// if fi.IsDir() && fj.IsFile() {
		// 	return true
		// } else if fi.IsFile() && fj.IsDir() {
		// 	return false
		// }
		return paw.ToLower(fi.BaseName) > paw.ToLower(fj.BaseName)
	}

	sortedFields = []string{"size", "modified", "accessed", "created", "name"}
	sortBy       = map[string]filetree.FilesBy{
		"size":             bySize,
		"modified":         byMTime,
		"accessed":         byATime,
		"created":          byCTime,
		"name":             byName,
		"reverse_size":     bySizeR,
		"reverse_modified": byMTimeR,
		"reverse_accessed": byATimeR,
		"reverse_created":  byCTimeR,
		"reverse_name":     byNameR,
	}
)

// TODO sortByField

func getSortOption(opt *gloption) *filetree.PDSortOption {
	if opt.isNoSort {
		return &filetree.PDSortOption{
			IsSort: !opt.isNoSort,
		}
	}

	// var sortOpt *filetree.PDSortOption
	if opt.isSortBySize {
		// paw.Info.Println("opt.isSortBySize", opt.isSortBySize)
		// paw.Info.Println("  opt.isReverse", opt.isReverse)
		if opt.isReverse {
			return &filetree.PDSortOption{
				IsSort:  true,
				SortWay: filetree.PDSortByReverseSize,
			}
		} else {
			return &filetree.PDSortOption{
				IsSort:  true,
				SortWay: filetree.PDSortBySize,
			}
		}
	}
	if opt.isSortByMTime {
		// paw.Info.Println("opt.isSortByMTime", opt.isSortByMTime)
		// paw.Info.Println("  opt.isReverse", opt.isReverse)
		if opt.isReverse {
			return &filetree.PDSortOption{
				IsSort:  true,
				SortWay: filetree.PDSortByReverseMtime,
			}
		} else {
			return &filetree.PDSortOption{
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
		return &filetree.PDSortOption{
			IsSort:  true,
			SortWay: filetree.PDSortByReverseName,
		}
	} else {
		return &filetree.PDSortOption{
			IsSort:  true,
			SortWay: filetree.PDSortByName,
		}
	}
	// }

	return nil
}
