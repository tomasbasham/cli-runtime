package iooption

import (
	"io"
	"io/fs"
	"os"
)

// OpenFile opens a file for reading. If the filename is "-" or empty then the
// standard input is returned. This is useful for reading from standard input
// when the filename is provided as a command line argument.
func OpenFile(fsys fs.FS, filename string) (io.ReadCloser, error) {
	if filename == "-" || filename == "" {
		return os.Stdin, nil
	}

	return fsys.Open(filename)
}
