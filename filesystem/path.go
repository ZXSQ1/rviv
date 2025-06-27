package filesystem

type Path struct {
	Filename string
	Filesys  Filesystem
}

type Paths struct {
	Filenames []string
	Filesys   Filesystem
}
