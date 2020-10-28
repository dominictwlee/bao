package fs

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	DirPerm = 0775
)

func CopyDir(srcdir, destdir string) (err error) {
	if err = os.MkdirAll(destdir, DirPerm); err != nil {
		return
	}

	entries, err := ioutil.ReadDir(srcdir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcdir, entry.Name())
		destPath := filepath.Join(destdir, entry.Name())

		fmt.Println(srcPath)
		fmt.Println(destPath)

		if entry.IsDir() {
			if err = CopyDir(srcPath, destPath); err != nil {
				return
			}
		} else {
			if err = CopyFile(srcPath, destPath); err != nil {
				return
			}
		}
	}

	return
}

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
