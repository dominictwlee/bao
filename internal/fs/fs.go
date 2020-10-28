package fs

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dest string) (err error) {
	srcStat, err := os.Stat(src)
	if err != nil {
		return
	}

	if !srcStat.Mode().IsRegular() {
		fmt.Errorf("%s is not a regular file", src)
	}

	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	if err = out.Sync(); err != nil {
		return
	}

	return
}
