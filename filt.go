package main

import (
	"github.com/shyang107/paw/filetree"
)

func getFiltOption(opt *gloption) *filetree.PDFilterOption {

	filtOpt := &filetree.PDFilterOption{
		IsFilt: false,
	}

	if opt.isJustFiles &&
		opt.isJustDirs { // -F -D
		info("Filt: files adn dirs")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFiles
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		!opt.isJustFiles { // -O
		info("Filt: no empty dir")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltNoEmptyDir
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		opt.isJustDirs &&
		!opt.isJustFiles { // -D -O
		info("Filt: just dirs, but no empty")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustDirsButNoEmpty
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		opt.isJustFiles { // -F -O
		info("Filt: just files, but no empty dir")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFilesButNoEmptyDir
		return filtOpt
	}

	if !opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		opt.isJustFiles { // -F
		info("Filt: just files")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFiles
		return filtOpt
	}

	if !opt.isNoEmptyDirs &&
		opt.isJustDirs &&
		!opt.isJustFiles { // -D
		info("Filt: just dirs")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustDirs
		return filtOpt
	}

	return filtOpt
}
