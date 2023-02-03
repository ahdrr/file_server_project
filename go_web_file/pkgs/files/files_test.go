package files

import (
	"filrserver/pkgs/model"
	"fmt"
	"io/fs"
	"path"
	"testing"
)

func TestIterDirectory(t *testing.T) {
	type Parse_request_struct struct {
		Pathurl   string
		Real_path string
		Statinfo  fs.FileInfo
	}
	tests := []struct {
		name string
	}{
		{"asd"}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := "/home/git_work/file_server_project/basedir/go_tools"
			fmt.Println(path.Base(filename))

		})
	}
}

func TestMvproject(t *testing.T) {

	tests := []struct {
		name string
		want []model.Mverr
	}{
		{}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mvproject("/tmp/",
				&model.Parse_request_mv{Filelist: []string{"/test"}, Dstdir: "/asd"})
			fmt.Println(got)
			//cc := got[0].err.(*os.LinkError).Err.Error()

		})
	}
}
