package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/shyang107/paw"
	"github.com/shyang107/paw/filetree"
)

type patflag int

const (
	allFlag patflag = 1 << iota
	includeFlag
	excludeFlag
	allincludeFlag      = allFlag | includeFlag
	allexcludeFlag      = allFlag | excludeFlag
	allinAndexcludeFlag = allFlag | includeFlag | excludeFlag
	inAndexcludeFlag    = includeFlag | excludeFlag
)

func getpatflag(opt *gloption) (pflag patflag) {
	if opt.isAllFiles && len(opt.excludePattern) == 0 && len(opt.includePattern) == 0 {
		pflag = allFlag
		goto END
	}
	if opt.isAllFiles && len(opt.excludePattern) > 0 && len(opt.includePattern) == 0 {
		pflag = allexcludeFlag
		goto END
	}
	if opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) > 0 {
		pflag = allinAndexcludeFlag
		goto END
	}
	if opt.isAllFiles &&
		len(opt.excludePattern) == 0 &&
		len(opt.includePattern) > 0 {
		pflag = allincludeFlag
		goto END
	}

	if !opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) == 0 {
		pflag = excludeFlag
		goto END
	}
	if !opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) > 0 {
		pflag = inAndexcludeFlag
		goto END
	}
	if !opt.isAllFiles &&
		len(opt.excludePattern) == 0 &&
		len(opt.includePattern) > 0 {
		pflag = includeFlag
		goto END
	}
END:
	return pflag
}

func optAllInAndExclude(opt *gloption, pdopt *filetree.PrintDirOption) {
	ren, err := regexp.Compile(opt.includePattern)
	if err != nil {
		paw.Error.Printf("including pattern: %q, error: %v", ren.String(), err)
		os.Exit(1)
	}
	rex, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		paw.Error.Printf("excluding pattern: %q, error: %v", rex.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !ren.MatchString(f.BaseName) && rex.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		// if !f.IsDir() {
		// 	if !ren.MatchString(f.BaseName) && rex.MatchString(f.BaseName) {
		// 		return filetree.SkipThis
		// 	}
		// }
		return nil
	}
}
func optInAndExclude(opt *gloption, pdopt *filetree.PrintDirOption) {
	ren, err := regexp.Compile(opt.includePattern)
	if err != nil {
		paw.Error.Printf("including pattern: %q, error: %v", ren.String(), err)
		os.Exit(1)
	}
	rex, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		paw.Error.Printf("excluding pattern: %q, error: %v", rex.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		_, file := filepath.Split(f.Path)
		if strings.HasPrefix(file, ".") {
			return filetree.SkipThis
		}
		if !ren.MatchString(f.BaseName) && rex.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		// if !f.IsDir() {
		// 	if !ren.MatchString(f.BaseName) && rex.MatchString(f.BaseName) {
		// 		return filetree.SkipThis
		// 	}
		// }
		return nil
	}
}
func optAllInclude(opt *gloption, pdopt *filetree.PrintDirOption) {
	re, err := regexp.Compile(opt.includePattern)
	if err != nil {
		paw.Error.Printf("including pattern: %q, error: %v", re.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !re.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		// if !f.IsDir() {
		// 	if !re.MatchString(f.BaseName) {
		// 		return filetree.SkipThis
		// 	}
		// }
		return nil
	}
}
func optInclude(opt *gloption, pdopt *filetree.PrintDirOption) {
	re, err := regexp.Compile(opt.includePattern)
	if err != nil {
		paw.Error.Printf("including pattern: %q, error: %v", re.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		_, file := filepath.Split(f.Path)
		if strings.HasPrefix(file, ".") {
			return filetree.SkipThis
		}
		if !re.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		// if !f.IsDir() {
		// 	if !re.MatchString(f.BaseName) {
		// 		return filetree.SkipThis
		// 	}
		// }
		return nil
	}
}

func optAllExclude(opt *gloption, pdopt *filetree.PrintDirOption) {
	re, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		paw.Error.Printf("excluding pattern: %q, error: %v", re.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if re.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		// if !f.IsDir() {
		// 	if re.MatchString(f.BaseName) {
		// 		return filetree.SkipThis
		// 	}
		// }
		return nil
	}
}

func optExclude(opt *gloption, pdopt *filetree.PrintDirOption) {
	re, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		paw.Error.Printf("excluding pattern: %q, error: %v", re.String(), err)
		os.Exit(1)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		_, file := filepath.Split(f.Path)
		if strings.HasPrefix(file, ".") {
			return filetree.SkipThis
		}
		if re.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		// if !f.IsDir() {
		// 	if re.MatchString(f.BaseName) {
		// 		return filetree.SkipThis
		// 	}
		// }
		return nil
	}
}

func optAllFiles(opt *gloption, pdopt *filetree.PrintDirOption) {
	pdopt.Ignore = func(f *filetree.File, e error) error {
		return nil
	}
}
