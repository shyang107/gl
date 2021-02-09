package main

import (
	"github.com/shyang107/paw/filetree"
)

func getFieldFlag(opt *gloption) filetree.PDFieldFlag {
	var (
		flag   filetree.PDFieldFlag
		fields []string
	)

	flag = filetree.PFieldPermissions
	fields = append(fields, "Permission")
	if opt.isFieldINode {
		flag = flag | filetree.PFieldINode
		fields = append(fields, "inode")
	}

	// if opt.isFieldPermissions {
	// 	flag = flag | filetree.PFieldPermissions
	// }

	if opt.isFieldLinks {
		flag = flag | filetree.PFieldLinks
		fields = append(fields, "Links")
	}
	// if opt.isFieldSize {
	// 	flag = flag | filetree.PFieldSize
	// }
	flag = flag | filetree.PFieldSize
	fields = append(fields, "Size")

	if opt.isFieldBlocks {
		flag = flag | filetree.PFieldBlocks
		fields = append(fields, "Blocks")
	}
	// if opt.isFieldUser {
	// 	flag = flag | filetree.PFieldUser
	// }
	flag = flag | filetree.PFieldUser
	fields = append(fields, "User")
	// if opt.isFieldGroup {
	// 	flag = flag | filetree.PFieldGroup
	// }
	flag = flag | filetree.PFieldGroup
	fields = append(fields, "Group")

	if opt.isFieldModified {
		flag = flag | filetree.PFieldModified
		fields = append(fields, "Modified")
	}
	if opt.isFieldAccessed {
		flag = flag | filetree.PFieldAccessed
		fields = append(fields, "Accessed")
	}
	if opt.isFieldCreated {
		flag = flag | filetree.PFieldCreated
		fields = append(fields, "Created")
	}
	if !opt.isFieldModified &&
		!opt.isFieldAccessed &&
		!opt.isFieldCreated {
		flag = flag | filetree.PFieldModified
		fields = append(fields, "Modified")
	}

	if opt.isFieldGit {
		flag = flag | filetree.PFieldGit
		fields = append(fields, "Git")
	}
	fields = append(fields, "Name")
	info("[getFieldFlag] fields includes %#v", fields)

	return flag
}
