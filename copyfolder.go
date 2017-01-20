package copyfolder

// FilePathsGetter takes a filepath and returns filenpaths in a slice
type FilePathsGetter interface {
	Get() ([]string, error)
}

// DstMaker returns a dst string
type DstMaker interface {
	MakeDst(string) (string, error)
}

// Copier copies from source to dest
type Copier interface {
	Copy(string, string) error
}

// FolderCopier interface which is used for copyfolder
type FolderCopier interface {
	FilePathsGetter
	DstMaker
	Copier
}

// Copy everything by usind the FolderCopier
func Copy(fc FolderCopier) error {
	files, err := fc.Get()
	if err != nil {
		return err
	}
	for _, f := range files {
		dst, err := fc.MakeDst(f)
		if err != nil {
			return err
		}
		err = fc.Copy(f, dst)
		if err != nil {
			return err
		}
	}
	return nil
}
