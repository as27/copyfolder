package fs

import (
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/as27/copyfolder"
)

// FolderCopier implements the copyfolder interface
type FolderCopier struct {
	srcRoot string
	dstRoot string
}

func NewFolderCopier(src, dst string) *FolderCopier {
	return &FolderCopier{
		srcRoot: src,
		dstRoot: dst,
	}
}

// Get returns all files, which can be found under srcRoot
func (fc *FolderCopier) Get() ([]string, error) {
	var filepaths []string
	filepath.Walk(
		fc.srcRoot,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				filepaths = append(filepaths, path)
			}
			return err
		},
	)
	return filepaths, nil
}

// MakeDst creates a dst path for a source
func (fc *FolderCopier) MakeDst(fpath string) (string, error) {
	relPath, err := filepath.Rel(fc.srcRoot, fpath)
	return path.Join(fc.dstRoot, relPath), err
}

// Copy copies a file
func (fc *FolderCopier) Copy(src, dst string) error {
	err := os.MkdirAll(filepath.Dir(dst), 0777)
	if err != nil {
		return err
	}
	fdst, err := os.Create(dst)
	defer fdst.Close()
	if err != nil {
		return err
	}

	fsrc, err := os.Open(src)
	defer fsrc.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(fdst, fsrc)
	if err != nil {
		return err
	}
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	mTime := srcInfo.ModTime()
	err = os.Chtimes(dst, mTime, mTime)
	if err != nil {
		return err
	}
	return nil
}

var a copyfolder.FolderCopier = &FolderCopier{}
