package main

import (
	"github.com/shyang107/paw/filetree"
)

func getFiltOption(opt *gloption) *filetree.PrintDirFilterOption {

	filtOpt := &filetree.PrintDirFilterOption{
		IsFilt: false,
	}

	if opt.isJustFiles &&
		opt.isJustDirs { // -F -D
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFiles
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		!opt.isJustFiles { // -O
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltNoEmptyDir
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		opt.isJustDirs &&
		!opt.isJustFiles { // -D -O
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustDirsButNoEmpty
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		opt.isJustFiles { // -F -O
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFilesButNoEmptyDir
		return filtOpt
	}

	if !opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		opt.isJustFiles { // -F
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFiles
		return filtOpt
	}

	if !opt.isNoEmptyDirs &&
		opt.isJustDirs &&
		!opt.isJustFiles { // -D
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustDirs
		return filtOpt
	}

	return filtOpt
}
