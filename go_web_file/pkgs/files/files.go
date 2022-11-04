package files

import (
	"filrserver/pkgs/model"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"github.com/h2non/filetype"
)

func IterDirectory(parse_request model.Parse_request_struct) (dirs model.Dir_stat, err error) {
	var f model.File
	var ftype string
	dirPath := parse_request.Real_path
	sinfo := parse_request.Statinfo
	dirs = model.Dir_stat{
		Dmsg: model.Dmsg{
			Code: 200,
			Msg:  "ok",
		},
		Data: []model.File{},
	}
	if sinfo.IsDir() {
		ldirs, err := os.ReadDir(dirPath)
		for _, dinfo := range ldirs {
			finfo, _ := dinfo.Info()
			fname := finfo.Name()
			if dinfo.IsDir() {
				ftype = "dir"
			} else {
				ftype = getfiletype(dirPath + fname)
			}
			f = getfileDetails(finfo)
			f.Filetype = ftype
			dirs.Data = append(dirs.Data, f)
		}
		return dirs, err
	}
	f = getfileDetails(sinfo)
	f.Filetype = getfiletype(dirPath)
	dirs.Data = append(dirs.Data, f)
	return dirs, nil
}

func getfileDetails(finfo fs.FileInfo) model.File {
	return model.File{
		Name: finfo.Name(),
		Time: finfo.ModTime().Format("2006-01-02 15:04:05"),
		Size: GetReadableFileSizeString(finfo.Size()),
	}
}

func GetReadableFileSizeString(fileSizeInBytes int64) string {
	var i = 0
	var fileSizefot = float64(fileSizeInBytes)
	byteUnits := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	for fileSizefot > 1024 {
		fileSizefot = fileSizefot / 1024
		i = i + 1
	}
	k := strconv.FormatFloat(fileSizefot, 'f', 2, 64)
	return k + byteUnits[i]
}

var f_buffer []byte = make([]byte, 261)

func getfiletype(filepath string) string {
	f, _ := os.Open(filepath)
	defer f.Close()
	n, _ := f.Read(f_buffer)
	contentType, _ := filetype.Match(f_buffer[0:n])
	if contentType == filetype.Unknown {
		fext := strings.Split(filepath, ".")
		if len(fext) > 1 {
			return fext[len(fext)-1]
		}
		return "file"
	}
	return contentType.Extension

}
