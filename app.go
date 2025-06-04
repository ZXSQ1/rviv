package main

import (
	"fmt"

	"github.com/ZXSQ1/rviv/ftpfs"
	"github.com/ZXSQ1/rviv/localfs"
	"github.com/ZXSQ1/rviv/processes"
)

func main() {
	ftpfilesys, _ := ftpfs.Connect("192.168.1.7", 2121, "omar", "dlpob")
	localfilesys := &localfs.LocalFS{}

	src := &processes.Path{
		Filename: "A1-Omar/Backups/documents[20250520].tar.xz",
		Filesys:  ftpfilesys,
	}

	dest := &processes.Path{
		Filename: "/home/omar/Temp/documents.tar.xz",
		Filesys:  localfilesys,
	}

	progress := make(chan int)

	go func() {
		fmt.Println(processes.Copy(src, dest, progress))
	}()

	for prog, ok := <-progress; ok; {
		fmt.Println(prog)
	}
}
