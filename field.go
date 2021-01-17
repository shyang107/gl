package main

import (
	"github.com/shyang107/paw/filetree"
)

func getFieldFlag(opt *gloption) filetree.PDFieldFlag {
	var flag filetree.PDFieldFlag

	flag = filetree.PFieldPermissions

	if opt.isFieldINode {
		flag = flag | filetree.PFieldINode
	}

	// if opt.isFieldPermissions {
	// 	flag = flag | filetree.PFieldPermissions
	// }

	if opt.isFieldLinks {
		flag = flag | filetree.PFieldLinks
	}
	// if opt.isFieldSize {
	// 	flag = flag | filetree.PFieldSize
	// }
	flag = flag | filetree.PFieldSize

	if opt.isFieldBlocks {
		flag = flag | filetree.PFieldBlocks
	}
	// if opt.isFieldUser {
	// 	flag = flag | filetree.PFieldUser
	// }
	flag = flag | filetree.PFieldUser
	// if opt.isFieldGroup {
	// 	flag = flag | filetree.PFieldGroup
	// }
	flag = flag | filetree.PFieldGroup

	if opt.isFieldModified {
		flag = flag | filetree.PFieldModified
	}
	if opt.isFieldAccessed {
		flag = flag | filetree.PFieldAccessed
	}
	if opt.isFieldCreated {
		flag = flag | filetree.PFieldCreated
	}
	if !opt.isFieldModified &&
		!opt.isFieldAccessed &&
		!opt.isFieldCreated {
		flag = flag | filetree.PFieldModified
	}

	if opt.isFieldGit {
		flag = flag | filetree.PFieldGit
	}

	return flag
}
