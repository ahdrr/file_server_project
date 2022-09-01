package model

import "io/fs"

type Parse_request_struct struct {
	Pathurl   string
	Real_path string
	Statinfo  fs.FileInfo
}

type Parse_request_rename struct {
	Newname string `yaml:"newname"`
}
