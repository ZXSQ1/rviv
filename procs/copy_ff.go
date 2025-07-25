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

	if err := CheckIsRegular(srcfile); err != nil {
		return err
	}

	srcStat, err := CheckStat(srcfile)

	if err != nil {
		return err
	}

	srcfileSize := srcStat.Size()

	if err := CheckExistsCreate(destfile); err != nil {
		return err
	}

	if err := CheckIsRegular(destfile); err != nil {
		return err
	}

	srcObj, err := CheckOpen(srcfile, filesystem.ModeRead)

	if err != nil {
		return err
	}

	defer srcObj.Close()

	destObj, err := CheckOpen(destfile, filesystem.ModeWrite)

	if err != nil {
		return err
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
			"unable to copy source file '%s' to destination file '%s'",
			ShowPath(srcfile), ShowPath(destfile),
		)
	}

	return nil
}
