package fs

import "testing"

func TestFolderCopier_Copy(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		fc      *FolderCopier
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fc.Copy(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("FolderCopier.Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
