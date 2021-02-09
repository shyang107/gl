package main

import (
	"path/filepath"
	"regexp"
	"strings"

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
	var ri, rx *regexp.Regexp
	if len(opt.includePattern) > 0 {
		ri, _ = regexp.Compile(opt.includePattern)
	}
	if len(opt.excludePattern) > 0 {
		rx, _ = regexp.Compile(opt.excludePattern)
	}

	if opt.isAllFiles &&
		len(opt.excludePattern) == 0 &&
		len(opt.includePattern) == 0 {
		info("[getpatflag] pattern: all")
		pflag = allFlag
		goto END
	}
	if opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) == 0 {
		info("[getpatflag] pattern: all; but excluding (regex: %q)", rx.String())
		pflag = allexcludeFlag
		goto END
	}
	if opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) > 0 {
		info("[getpatflag] pattern: all; but including (regex: %q) and excluding (regex: %q)", ri.String(), rx.String())
		pflag = allinAndexcludeFlag
		goto END
	}
	if opt.isAllFiles &&
		len(opt.excludePattern) == 0 &&
		len(opt.includePattern) > 0 {
		info("[getpatflag] pattern: all; but including (regex: %q)", ri.String())
		pflag = allincludeFlag
		goto END
	}

	if !opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) == 0 {
		info("[getpatflag] pattern: excluding (regex: %q)", rx.String())
		pflag = excludeFlag
		goto END
	}
	if !opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) > 0 {
		info("[getpatflag] pattern: including (regex: %q) and excluding (regex: %q)", ri.String(), rx.String())
		pflag = inAndexcludeFlag
		goto END
	}
	if !opt.isAllFiles &&
		len(opt.excludePattern) == 0 &&
		len(opt.includePattern) > 0 {
		info("[getpatflag] pattern: including (regex: %q)", ri.String())
		pflag = includeFlag
		goto END
	}
END:
	return pflag
}

func optAllInAndExclude(opt *gloption, pdopt *filetree.PrintDirOption) {
	ri, err := regexp.Compile(opt.includePattern)
	if err != nil {
		fatal("[optAllInAndExclude] including pattern: %q, error: %v", ri.String(), err)
	}
	rx, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		fatal("excluding pattern: %q, error: %v", rx.String(), err)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !ri.MatchString(f.BaseName) && rx.MatchString(f.BaseName) {
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
	ri, err := regexp.Compile(opt.includePattern)
	if err != nil {
		fatal("including pattern: %q, error: %v", ri.String(), err)
	}
	rx, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		fatal("excluding pattern: %q, error: %v", rx.String(), err)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		_, file := filepath.Split(f.Path)
		if strings.HasPrefix(file, ".") {
			return filetree.SkipThis
		}
		if !ri.MatchString(f.BaseName) && rx.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		// if !f.IsDir() {
		// 	if !ri.MatchString(f.BaseName) && rex.MatchString(f.BaseName) {
		// 		return filetree.SkipThis
		// 	}
		// }
		return nil
	}
}
func optAllInclude(opt *gloption, pdopt *filetree.PrintDirOption) {
	ri, err := regexp.Compile(opt.includePattern)
	if err != nil {
		fatal("including pattern: %q, error: %v", ri.String(), err)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !ri.MatchString(f.BaseName) {
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
	ri, err := regexp.Compile(opt.includePattern)
	if err != nil {
		fatal("including pattern: %q, error: %v", ri.String(), err)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		_, file := filepath.Split(f.Path)
		if strings.HasPrefix(file, ".") {
			return filetree.SkipThis
		}
		if !ri.MatchString(f.BaseName) {
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
	rx, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		fatal("excluding pattern: %q, error: %v", rx.String(), err)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if rx.MatchString(f.BaseName) {
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
	rx, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		fatal("excluding pattern: %q, error: %v", rx.String(), err)
	}
	pdopt.Ignore = func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		_, file := filepath.Split(f.Path)
		if strings.HasPrefix(file, ".") {
			return filetree.SkipThis
		}
		if rx.MatchString(f.BaseName) {
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
