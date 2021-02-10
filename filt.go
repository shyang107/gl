package main

import (
	"github.com/shyang107/paw/filetree"
	"github.com/sirupsen/logrus"
)

func getFiltOption(opt *gloption) *filetree.PDFilterOption {

	filtOpt := &filetree.PDFilterOption{
		IsFilt: false,
	}

	var fields = logrus.Fields{
		"isNoEmptyDirs": opt.isNoEmptyDirs,
		"isJustDirs":    opt.isJustDirs,
		"isJustFiles":   opt.isJustFiles,
	}
	if opt.isJustFiles &&
		opt.isJustDirs { // -F -D
		lg.WithFields(fields).Trace("files adn dirs")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFiles
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		!opt.isJustFiles { // -O
		lg.WithFields(fields).Trace("no empty dir")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltNoEmptyDir
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		opt.isJustDirs &&
		!opt.isJustFiles { // -D -O
		lg.WithFields(fields).Trace("just dirs, but no empty")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustDirsButNoEmpty
		return filtOpt
	}

	if opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		opt.isJustFiles { // -F -O
		lg.WithFields(fields).Trace("just files, but no empty dir")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFilesButNoEmptyDir
		return filtOpt
	}

	if !opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		opt.isJustFiles { // -F
		lg.WithFields(fields).Trace("just files")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFiles
		return filtOpt
	}

	if !opt.isNoEmptyDirs &&
		opt.isJustDirs &&
		!opt.isJustFiles { // -D
		lg.WithFields(fields).Trace("just dirs")
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustDirs
		return filtOpt
	}

	return filtOpt
}
