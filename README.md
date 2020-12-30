# gl

`gl path` list files of path in color view.

`gl` is in spired by [macOS · exa](https://the.exa.website/install/macos).

```none
NAME:
   gl - list directory (excluding hidden items) in color view.

USAGE:
   gl [global options] command [command options] [path]

VERSION:
   0.0.2

AUTHOR:
   Shuhhua Yang <shyang107@gmail.com>

COMMANDS:
   version, v, V  print only the version
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --list, -l                     print out in list view (default: true)
   --listtree, -t                 print out in the view of combining list and tree (default: false)
   --tree, -T                     print out in the tree view (default: false)
   --table, -b                    print out in the table view (default: false)
   --level, -L                    print out in the level view (default: false)
   --classify, -F                 display type indicator by file names (default: false)
   --depth value, -d value        print out in the level view (default: 0)
   --recurse, -R                  recurse into directories (equivalent to --depth=-1) (default: false)
   --all, -a                      show all file including hidden files (default: false)
   --include pattern, -n pattern  set regex pattern to include some files, applied to file only
   --exclude pattern, -x pattern  set regex pattern to exclude some files, applied to file only
   --help, -h                     show help (default: false)
   --version, -v, -V              print only the version (default: false)
```
