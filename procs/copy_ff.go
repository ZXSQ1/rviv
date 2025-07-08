package procs

import (
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/info"
	"github.com/schollz/progressbar/v3"
)

func CopyFile(srcfile, destfile config.Path, verbose bool,
	printfn func(src, dest string)) error {

	srcfileSize := 0

	if !srcfile.Fsys.IsExist(srcfile.Filename) {
		return info.Error("source file '%s' does not exist", ShowPath(srcfile))
	} else {
		stat, err := srcfile.Fsys.Stat(srcfile.Filename)

		if err != nil {
			return info.Error("unable to stat '%s'", ShowPath(srcfile))
		}

		srcfileSize = int(stat.Size())

		if !stat.Mode().IsRegular() {
			return info.Error(
				"source file '%s' is not a regular file", ShowPath(srcfile),
			)
		}
	}

	if !destfile.Fsys.IsExist(destfile.Filename) {
		if destfile.Fsys.Create(destfile.Filename) != nil {
			return info.Error(
				"unable to create file '%s'", ShowPath(destfile),
			)
		}
	} else {
		stat, err := destfile.Fsys.Stat(destfile.Filename)

		if err != nil {
			return info.Error(
				"unable to stat '%s'", ShowPath(destfile),
			)
		}

		if !stat.Mode().IsRegular() {
			return info.Error(
				"destination file '%s' is not a regular file",
				ShowPath(destfile),
			)
		}
	}

	srcObj, err := srcfile.Fsys.Open(srcfile.Filename, filesystem.ModeRead)

	if err != nil {
		return info.Error("unable to open '%s' for reading", ShowPath(srcfile))
	}

	defer srcObj.Close()
	destObj, err := destfile.Fsys.Open(destfile.Filename,
		filesystem.ModeWrite)

	if err != nil {
		return info.Error("unable to open '%s' for writing", ShowPath(destfile))
	}

	defer destObj.Close()

	multiWriter := io.MultiWriter(destObj)

	if verbose {
		printfn(
			ShowPath(srcfile),
			ShowPath(destfile),
		)

		bar := progressbar.NewOptions(
			srcfileSize,
			progressbar.OptionSetDescription(info.Padding+info.Padding),
			progressbar.OptionShowBytes(true),
			progressbar.OptionClearOnFinish(),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        ".",
				SaucerHead:    ".",
				SaucerPadding: "",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)

		multiWriter = io.MultiWriter(destObj, bar)
	}

	_, err = io.Copy(multiWriter, srcObj)

	if err != nil {
		return info.Error(
			"unable to copy source file '%s' to destination file '%s'", ShowPath(
				srcfile), ShowPath(destfile),
		)
	}

	return nil
}
