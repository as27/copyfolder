package copyfolder

// FolderCopier interface which is used for copyfolder
type FolderCopier interface {
	Get() ([]string, error)
	MakeDst(string) (string, error)
	Copy(string, string) error
}

// Copy everything by using the FolderCopier
func Copy(fc FolderCopier) []error {
	var errs []error
	files, err := fc.Get()
	if err != nil {
		errs = append(errs, err)
	}
	for _, f := range files {
		dst, err := fc.MakeDst(f)
		if err != nil {
			errs = append(errs, err)
		}
		err = fc.Copy(f, dst)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
