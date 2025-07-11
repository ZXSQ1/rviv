package procs

import (
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/info"
	"github.com/schollz/progressbar/v3"
)

func MoveFile(srcfile, destfile config.Path, verbose bool,
	printfn func(src, dest string)) error {

	if err := AssertExists(srcfile); err != nil {
		return err
	}

	if err := AssertIsRegular(srcfile); err != nil {
		return err
	}

	srcStat, err := AssertStatWorks(srcfile)

	if err != nil {
		return err
	}

	srcfileSize := srcStat.Size()

	if err := AssertExistsCreate(destfile); err != nil {
		return err
	}

	if err := AssertIsRegular(destfile); err != nil {
		return err
	}

	srcObj, err := srcfile.Fsys.Open(srcfile.Filename, filesystem.ModeRead)

	if err != nil {
		return info.Error("unable to open '%s' for reading", ShowPath(srcfile))
	}

	defer srcObj.Close()
	destObj, err := destfile.Fsys.Open(
		destfile.Filename, filesystem.ModeWrite,
	)

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
			int(srcfileSize),

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
			"unable to move source file '%s' to destination file '%s'",
			ShowPath(srcfile), ShowPath(destfile),
		)
	}

	if srcObj.Close() != nil {
		return info.Error("unable to close file '%s'", ShowPath(srcfile))
	}

	if srcfile.Fsys.Remove(srcfile.Filename) != nil {
		return info.Error("unable to remove file '%s'", ShowPath(srcfile))
	}

	return nil
}
