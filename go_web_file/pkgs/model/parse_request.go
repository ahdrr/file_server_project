package model

import (
	"io/fs"
)

type Parse_request_struct struct {
	Pathurl   string
	Real_path string
	Statinfo  fs.FileInfo
}

type Parse_request_rename struct {
	Newname string `json:"newname" binding:"required"`
}

type Parse_request_list struct {
	Onlydir bool `json:"onlydir"`
}

type Parse_request_mv struct {
	Filelist []string `json:"filelist" binding:"required,min=1"`
	Dstdir   string   `json:"dstdir" binding:"required"`
}
