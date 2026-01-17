package iooption_test

import (
	"testing"
	"testing/fstest"

	"github.com/tomasbasham/cli-runtime/iooption"
)

var testdata = fstest.MapFS{
	"testdata/test_file.txt": &fstest.MapFile{
		Data: []byte("This is a test file.\nIt contains sample text for testing purposes."),
	},
}

func TestOpenFile(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path    string
		wantErr bool
	}{
		"open file": {
			path: "testdata/test_file.txt",
		},
		"open stdin": {
			path: "-",
		},
		"returns error for non-existent file": {
			path:    "testdata/non_existent.txt",
			wantErr: true,
		},
		"returns error for directory path": {
			path:    "testdata/",
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			f, err := iooption.OpenFile(testdata, tt.path)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if f != nil {
				defer f.Close()
			}
			if !tt.wantErr && f == nil {
				t.Errorf("expected non-nil file handle")
			}
		})
	}
}
