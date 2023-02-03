package files

import (
	"filrserver/pkgs/model"
	"filrserver/pkgs/zlog"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/h2non/filetype"
)

func IterDirectory(parse_request model.Parse_request_struct, odir bool) (dirs model.Dir_stat, err error) {
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
				if odir {
					continue
				}
				ftype = getfiletype(dirPath + fname)
			}
			f = getfileDetails(finfo)
			f.Filetype = ftype
			dirs.Data = append(dirs.Data, f)
		}
		return dirs, err
	}
	if odir {
		return dirs, nil
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

func getfiletype(filepath string) string {
	var f_buffer []byte = make([]byte, 261)
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

func FileErrPars(e error) string {
	if e == nil {
		return ""
	}
	switch e.(type) {
	case *os.LinkError:
		return e.(*os.LinkError).Err.Error()
	case *fs.PathError:
		return e.(*fs.PathError).Err.Error()
	default:
		zlog.SugLog.Warnf("错误类型未断言: %v", e)
		return "other error"
	}

}

func Mvproject(real_path string, request_mv *model.Parse_request_mv) model.MvResoponse {
	var wg sync.WaitGroup
	var rg sync.WaitGroup

	var errs *[]model.Mverr = new([]model.Mverr)
	var responseChannel = make(chan model.Mverr, 10)
	var errNums = make(chan int, 1)
	for _, s := range request_mv.Filelist {
		wg.Add(1)
		go func(s string, respchan chan model.Mverr) {
			spath := filepath.Join(real_path, s)
			sbase := filepath.Base(spath)
			newfile := filepath.Join(real_path, request_mv.Dstdir, sbase)
			err := os.Rename(spath, newfile)
			respchan <- model.Mverr{Name: s, Err: FileErrPars(err)}
			wg.Done()
		}(s, responseChannel)
	}
	go func(errs *[]model.Mverr, respchan chan model.Mverr, errNums chan int) {
		rg.Add(1)
		var x int
		for rc := range responseChannel {
			if rc.Err != "" {
				x += 1
			}
			*errs = append(*errs, rc)
		}
		errNums <- x
		rg.Done()
	}(errs, responseChannel, errNums)
	wg.Wait()
	close(responseChannel)
	rg.Wait()
	em := <-errNums
	close(errNums)
	return model.MvResoponse{
		Dmsg: model.Dmsg{
			Msg: fmt.Sprintf("mv done,%v success,%v failed",
				len(request_mv.Filelist)-em,
				em)},
		Err: *errs,
	}

}
