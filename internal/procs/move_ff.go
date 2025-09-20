package procs

import (
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/info"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

func MoveFile(srcfile, destfile config.Path, verbose bool,
	printfn func(src, dest string)) error {

	if err := IsRegular(srcfile); err != nil {
		return err
	}

	srcStat, err := Stat(srcfile)

	if err != nil {
		return err
	}

	if err := IsExist(destfile, CreateFile); err != nil {
		return err
	}

	if err := IsRegular(destfile); err != nil {
		return err
	}

	srcObj, err := Open(srcfile, filesystem.ModeRead)

	if err != nil {
		return err
	}

	destObj, err := Open(destfile, filesystem.ModeWrite)

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
			int(srcStat.Size()),

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

	if err != nil && !errors.Is(err, filesystem.ErrEOF) {
		return info.Error(
			"unable to move source file '%s' to destination file '%s'",
			ShowPath(srcfile), ShowPath(destfile),
		)
	}

	if srcObj.Close() != nil {
		return info.Error("unable to close file '%s'", ShowPath(srcfile))
	}

	if err := RemoveFile(srcfile); err != nil {
		return err
	}

	return nil
}
