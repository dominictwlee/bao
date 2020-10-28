package fs

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCopyFile(t *testing.T) {
	err := CopyFile("testfiles/dummy.txt", "dummy.txt")
	if err != nil {
		t.Error(err)
	}

	srcFile, err := ioutil.ReadFile("testfiles/dummy.txt")
	if err != nil {
		t.Error(err)
	}

	destFile, err := ioutil.ReadFile("dummy.txt")
	if err != nil {
		t.Error("file was copied to the wrong destination")
	}

	if string(srcFile) != string(destFile) {
		t.Error("Copied data is not the same")
	}

	cleanup()
}

func cleanup() {
	if err := os.Remove("dummy.txt"); err != nil {
		log.Fatalln(err)
	}
}
