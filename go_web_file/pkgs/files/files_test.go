package files

import (
	"filrserver/pkgs/model"
	"fmt"
	"io/fs"
	"os"
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
		{}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := "/home/git_work/file_server_project/basedir/go_tools"
			statinfo, _ := os.Stat("/home/git_work/file_server_project/basedir/go_tools")
			b, _ := IterDirectory(model.Parse_request_struct{
				Real_path: path,
				Statinfo:  statinfo,
			})
			fmt.Printf("%#v\n", b)
		})
	}
}
