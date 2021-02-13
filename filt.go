package main

import (
	"github.com/shyang107/paw/filetree"
	"github.com/sirupsen/logrus"
)

func getFiltOption(opt *gloption) *filetree.PDFilterOption {

	filtOpt := &filetree.PDFilterOption{
		IsFilt: false,
	}

	if opt.isJustFiles &&
		opt.isJustDirs { // -F -D
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFiles
		goto END
	}

	if opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		!opt.isJustFiles { // -O
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltNoEmptyDir
		goto END
	}

	if opt.isNoEmptyDirs &&
		opt.isJustDirs &&
		!opt.isJustFiles { // -D -O
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustDirsButNoEmpty
		goto END
	}

	if opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		opt.isJustFiles { // -F -O
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFilesButNoEmptyDir
		goto END
	}

	if !opt.isNoEmptyDirs &&
		!opt.isJustDirs &&
		opt.isJustFiles { // -F
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustFiles
		goto END
	}

	if !opt.isNoEmptyDirs &&
		opt.isJustDirs &&
		!opt.isJustFiles { // -D
		filtOpt.IsFilt = true
		filtOpt.FiltWay = filetree.PDFiltJustDirs
		goto END
	}

END:
	lg.WithFields(logrus.Fields{
		"IsFilt":  filtOpt.IsFilt,
		"FiltWay": filtOpt.FiltWay.String(),
	}).Trace("filt")

	return filtOpt
}
