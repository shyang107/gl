package main

import (
	"github.com/shyang107/paw/filetree"
	"github.com/urfave/cli"
)

type gloption struct {
	path           string
	isList         bool
	isListTree     bool
	isTree         bool
	isTable        bool
	isLevel        bool
	isClassify     bool
	isRecurse      bool
	depth          int
	isAllFiles     bool
	includePattern string
	excludePattern string

	isNoEmptyDirs bool
	isJustFiles   bool
	isJustDirs    bool

	isNoSort      bool
	isReverse     bool
	sortByField   string
	isSortByName  bool //default name
	isSortBySize  bool
	isSortByMTime bool
	isGrouped     bool
	isExtended    bool

	// view field
	isFieldINode       bool // inode
	isFieldPermissions bool // permissions
	isFieldLinks       bool
	isFieldSize        bool
	isFieldBlocks      bool
	isFieldUser        bool
	isFieldGroup       bool
	isFieldModified    bool // date modified
	isFieldAccessed    bool // date accessed
	isFieldCreated     bool // date created
	isFieldGit         bool
}

var (
	opt   = new(gloption)
	pdopt = filetree.NewPrintDirOption()
	err   error

	listFlag = cli.BoolFlag{
		Name:        "list",
		Aliases:     []string{"l"},
		Value:       true,
		Usage:       "print out in list view",
		Destination: &opt.isList,
	}
	listTreeFlag = cli.BoolFlag{
		Name:        "listtree",
		Aliases:     []string{"t"},
		Value:       false,
		Usage:       "print out in the view of combining list and tree",
		Destination: &opt.isListTree,
	}
	treeFlag = cli.BoolFlag{
		Name:        "tree",
		Aliases:     []string{"T"},
		Value:       false,
		Usage:       "print out in the tree view",
		Destination: &opt.isTree,
	}
	tableFlag = cli.BoolFlag{
		Name:        "table",
		Aliases:     []string{"b"},
		Value:       false,
		Usage:       "print out in the table view",
		Destination: &opt.isTable,
	}
	levelFlag = cli.BoolFlag{
		Name:        "level",
		Aliases:     []string{"L"},
		Value:       false,
		Usage:       "print out in the level view",
		Destination: &opt.isLevel,
	}
	clsassifyFlag = cli.BoolFlag{
		Name:        "classify",
		Aliases:     []string{"f"},
		Value:       false,
		Usage:       "display type indicator by file names",
		Destination: &opt.isClassify,
	}
	depthFlag = cli.IntFlag{
		Name:        "depth",
		Aliases:     []string{"d"},
		Value:       0,
		Usage:       "print out in the level view",
		Destination: &opt.depth,
	}
	recurseFlag = cli.BoolFlag{
		Name:        "recurse",
		Aliases:     []string{"R"},
		Value:       false,
		Usage:       "recurse into directories (equivalent to --depth=-1)",
		Destination: &opt.isRecurse,
	}
	allFilesFlag = cli.BoolFlag{
		Name:        "all",
		Aliases:     []string{"a"},
		Value:       false,
		Usage:       "show all files including hidden files",
		Destination: &opt.isAllFiles,
	}
	includePatternFlag = cli.StringFlag{
		Name:        "include",
		Aliases:     []string{"ri"},
		Value:       "",
		Usage:       "set regex `pattern` to include some files, applied to file only",
		Destination: &opt.includePattern,
	}
	excludePatternFlag = cli.StringFlag{
		Name:        "exclude",
		Aliases:     []string{"rx"},
		Value:       "",
		Usage:       "set regex `pattern` to exclude some files, applied to file only",
		Destination: &opt.excludePattern,
	}

	isNoEmptyDirsFlag = cli.BoolFlag{
		Name:        "no-empty-dirs",
		Aliases:     []string{"O"},
		Value:       false,
		Usage:       "show all files but not empty directories",
		Destination: &opt.isNoEmptyDirs,
	}
	isJustFilesFlag = cli.BoolFlag{
		Name:        "just-files",
		Aliases:     []string{"F"},
		Value:       false,
		Usage:       "show all files but not directories, has high priority than --just-dirs",
		Destination: &opt.isJustFiles,
	}
	isJustDirsFlag = cli.BoolFlag{
		Name:        "just-dirs",
		Aliases:     []string{"D"},
		Value:       false,
		Usage:       "show all dirs but not files",
		Destination: &opt.isJustDirs,
	}

	isNoSortFlag = cli.BoolFlag{
		Name:        "no-sort",
		Aliases:     []string{"N"},
		Value:       false,
		Usage:       "not sort by name in increasing order (single key)",
		Destination: &opt.isNoSort,
	}
	isReverseFlag = cli.BoolFlag{
		Name:        "reverse",
		Aliases:     []string{"r"},
		Value:       false,
		Usage:       "sort in decreasing order, default sort by name",
		Destination: &opt.isReverse,
	}
	sortByFieldFlag = cli.StringFlag{
		Name:        "sort",
		Aliases:     []string{"sf"},
		Value:       "name",
		Usage:       "which single `field` to sort by. (field: inode, links, block, size, mtime (ot modified), atime (or accessed), ctime (or created), name; «field»[R]: reverse sort)",
		Destination: &opt.sortByField,
	}
	isSortByNameFlag = cli.BoolFlag{
		Name:        "sort-by-name",
		Aliases:     []string{"sn"},
		Value:       false,
		Usage:       "sort by name in increasing order (single key)",
		Destination: &opt.isSortByName,
	}
	isSortBySizeFlag = cli.BoolFlag{
		Name:        "sort-by-size",
		Aliases:     []string{"sz"},
		Value:       false,
		Usage:       "sort by size in increasing order (single key)",
		Destination: &opt.isSortBySize,
	}
	isSortByMTimeFlag = cli.BoolFlag{
		Name:        "sort-by-mtime",
		Aliases:     []string{"sm"},
		Value:       false,
		Usage:       "sort by modified time in increasing order (single key)",
		Destination: &opt.isSortByMTime,
	}

	isGroupedFlag = cli.BoolFlag{
		Name:        "grouped",
		Aliases:     []string{"g"},
		Value:       false,
		Usage:       "group files and directories separately",
		Destination: &opt.isGrouped,
	}

	isExtendedFlag = cli.BoolFlag{
		Name:        "extended",
		Aliases:     []string{"@"},
		Value:       false,
		Usage:       "list each file's extended attributes and sizes",
		Destination: &opt.isExtended,
	}

	isFieldINodeFlag = cli.BoolFlag{
		Name:        "inode",
		Aliases:     []string{"i"},
		Value:       false,
		Usage:       "list each file's inode number",
		Destination: &opt.isFieldINode,
	}
	// isFieldPermissionsFlag = cli.BoolFlag{
	// 	Name:        "permissions",
	// 	Aliases:     []string{"p"},
	// 	Value:       false,
	// 	Usage:       "list each file's permissions",
	// 	Destination: &opt.isFieldPermissions,
	// }
	isFieldLinksFlag = cli.BoolFlag{
		Name:        "links",
		Aliases:     []string{"H"},
		Value:       false,
		Usage:       "list each file's number of hard links",
		Destination: &opt.isFieldLinks,
	}
	// isFieldSizeFlag = cli.BoolFlag{
	// 	Name:        "size",
	// 	Aliases:     []string{"S"},
	// 	Value:       false,
	// 	Usage:       "list each file's size",
	// 	Destination: &opt.isFieldSize,
	// }
	isFieldBlocksFlag = cli.BoolFlag{
		Name:        "blocks",
		Aliases:     []string{"B"},
		Value:       false,
		Usage:       "show number of file system blocks",
		Destination: &opt.isFieldBlocks,
	}
	// isFieldUserFlag = cli.BoolFlag{
	// 	Name:        "user",
	// 	Aliases:     []string{"ur"},
	// 	Value:       false,
	// 	Usage:       "show user's name",
	// 	Destination: &opt.isFieldUser,
	// }
	// isFieldGroupFlag = cli.BoolFlag{
	// 	Name:        "group",
	// 	Aliases:     []string{"gp"},
	// 	Value:       false,
	// 	Usage:       "show user's group name",
	// 	Destination: &opt.isFieldGroup,
	// }
	isFieldGitFlag = cli.BoolFlag{
		Name: "git",
		// Aliases:     []string{"gp"},
		Value:       false,
		Usage:       " list each file's Git status, if tracked or ignored",
		Destination: &opt.isFieldGit,
	}

	isModifiedFlag = cli.BoolFlag{
		Name:        "modified",
		Aliases:     []string{"m"},
		Value:       false,
		Usage:       "use the modified timestamp field",
		Destination: &opt.isFieldModified,
	}
	isAccessedFlag = cli.BoolFlag{
		Name:        "accessed",
		Aliases:     []string{"u"},
		Value:       false,
		Usage:       "use the accessed timestamp field",
		Destination: &opt.isFieldAccessed,
	}
	isCreatedFlag = cli.BoolFlag{
		Name:        "created",
		Aliases:     []string{"U"},
		Value:       false,
		Usage:       "use the created timestamp field",
		Destination: &opt.isFieldCreated,
	}
)
