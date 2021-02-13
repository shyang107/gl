package main

import (
	"fmt"
	"regexp"

	"github.com/shyang107/paw/filetree"
	"github.com/sirupsen/logrus"
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

func (p patflag) Ignore(opt *gloption) filetree.IgnoreFunc {
	switch p {
	case allFlag:
		return optAllFiles(opt)
	case includeFlag:
		return optInclude(opt)
	case excludeFlag:
		return optExclude(opt)
	case allincludeFlag:
		return optAllInclude(opt) // allFlag | includeFlag
	case allexcludeFlag:
		return optAllExclude(opt) // allFlag | excludeFlag
	case allinAndexcludeFlag:
		return optAllInAndExclude(opt) // allFlag | includeFlag | excludeFlag
	case inAndexcludeFlag:
		return optInAndExclude(opt) // includeFlag | excludeFlag
	default:
		return filetree.DefaultIgnoreFn
	}
}

func getPatternflag(opt *gloption) (pflag patflag) {

	if opt.isAllFiles &&
		len(opt.excludePattern) == 0 &&
		len(opt.includePattern) == 0 {
		lg.WithFields(logrus.Fields{
			"ri_pat": fmt.Sprintf("%q", opt.includePattern),
			"rx_pat": fmt.Sprintf("%q", opt.excludePattern),
		}).Trace("pattern: all")
		pflag = allFlag
		goto END
	}
	if opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) == 0 {
		lg.WithFields(logrus.Fields{
			"ri_pat": fmt.Sprintf("%q", opt.includePattern),
			"rx_pat": fmt.Sprintf("%q", opt.excludePattern),
		}).Trace("pattern: all; but excluding")
		pflag = allexcludeFlag
		goto END
	}
	if opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) > 0 {
		lg.WithFields(logrus.Fields{
			"ri_pat": fmt.Sprintf("%q", opt.includePattern),
			"rx_pat": fmt.Sprintf("%q", opt.excludePattern),
		}).Trace("pattern: all; but including and excluding")
		pflag = allinAndexcludeFlag
		goto END
	}
	if opt.isAllFiles &&
		len(opt.excludePattern) == 0 &&
		len(opt.includePattern) > 0 {
		lg.WithFields(logrus.Fields{
			"ri_pat": fmt.Sprintf("%q", opt.includePattern),
			"rx_pat": fmt.Sprintf("%q", opt.excludePattern),
		}).Trace("pattern: all; but including")
		pflag = allincludeFlag
		goto END
	}

	if !opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) == 0 {
		lg.WithFields(logrus.Fields{
			"ri_pat": fmt.Sprintf("%q", opt.includePattern),
			"rx_pat": fmt.Sprintf("%q", opt.excludePattern),
		}).Trace("pattern: excluding")
		pflag = excludeFlag
		goto END
	}
	if !opt.isAllFiles &&
		len(opt.excludePattern) > 0 &&
		len(opt.includePattern) > 0 {
		lg.WithFields(logrus.Fields{
			"ri_pat": fmt.Sprintf("%q", opt.includePattern),
			"rx_pat": fmt.Sprintf("%q", opt.excludePattern),
		}).Trace("pattern: including and excluding")
		pflag = inAndexcludeFlag
		goto END
	}
	if !opt.isAllFiles &&
		len(opt.excludePattern) == 0 &&
		len(opt.includePattern) > 0 {
		lg.WithFields(logrus.Fields{
			"ri_pat": fmt.Sprintf("%q", opt.includePattern),
			"rx_pat": fmt.Sprintf("%q", opt.excludePattern),
		}).Trace("pattern: including")
		pflag = includeFlag
		goto END
	}
END:
	return pflag
}

func optAllInAndExclude(opt *gloption) filetree.IgnoreFunc {
	ri, err := regexp.Compile(opt.includePattern)
	if err != nil {
		fatal("[optAllInAndExclude] including pattern: %q: %v", ri.String(), err)
	}
	rx, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		fatal("excluding pattern: %q: %v", rx.String(), err)
	}
	return func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !ri.MatchString(f.BaseName) && rx.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		return nil
	}
}
func optInAndExclude(opt *gloption) filetree.IgnoreFunc {
	ri, err := regexp.Compile(opt.includePattern)
	if err != nil {
		fatal("including pattern: %q: %v", ri.String(), err)
	}
	rx, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		fatal("excluding pattern: %q: %v", rx.String(), err)
	}
	return func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if errg := filetree.DefaultIgnoreFn(f, nil); errg != nil {
			return filetree.SkipThis
		}
		if !ri.MatchString(f.BaseName) && rx.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		return nil
	}
}
func optAllInclude(opt *gloption) filetree.IgnoreFunc {
	ri, err := regexp.Compile(opt.includePattern)
	if err != nil {
		fatal("including pattern: %q: %v", ri.String(), err)
	}
	return func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if !ri.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		return nil
	}
}
func optInclude(opt *gloption) filetree.IgnoreFunc {
	ri, err := regexp.Compile(opt.includePattern)
	if err != nil {
		fatal("including pattern: %q: %v", ri.String(), err)
	}
	return func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if errg := filetree.DefaultIgnoreFn(f, nil); errg != nil {
			return filetree.SkipThis
		}
		if !ri.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		return nil
	}
}

func optAllExclude(opt *gloption) filetree.IgnoreFunc {
	rx, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		fatal("excluding pattern: %q: %v", rx.String(), err)
	}
	return func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if rx.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		return nil
	}
}

func optExclude(opt *gloption) filetree.IgnoreFunc {
	rx, err := regexp.Compile(opt.excludePattern)
	if err != nil {
		fatal("excluding pattern: %q: %v", rx.String(), err)
	}
	return func(f *filetree.File, e error) error {
		if err != nil {
			return err
		}
		if errg := filetree.DefaultIgnoreFn(f, nil); errg != nil {
			return filetree.SkipThis
		}
		if rx.MatchString(f.BaseName) {
			return filetree.SkipThis
		}
		return nil
	}
}

func optAllFiles(opt *gloption) filetree.IgnoreFunc {
	return func(f *filetree.File, e error) error {
		return nil
	}
}
