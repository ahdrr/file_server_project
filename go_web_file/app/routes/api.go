package routes

import (
	"filrserver/pkgs/config"
	"filrserver/pkgs/files"
	"filrserver/pkgs/model"
	"filrserver/pkgs/systeminfo"
	"os"

	//"filrserver/pkgs/zlog"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func get_realpath(c *gin.Context) (string, string) {
	basedir := config.ViperConfig.GetString("basedir")
	pathurl := c.Param("pathurl")
	real_path := filepath.Join(basedir, c.GetString("role"), pathurl)
	return pathurl, real_path
}

func check_pathurl(c *gin.Context) (parse_request model.Parse_request_struct, err error) {
	pathurl, real_path := get_realpath(c)
	statinfo, err := os.Stat(real_path)
	if err != nil {
		if os.IsNotExist(err) {
			c.AbortWithStatusJSON(500, gin.H{
				"code": 2003,
				"msg":  "{{basedir}}" + pathurl + " not exits",
			},
			)
		} else {

			c.AbortWithStatusJSON(500, gin.H{
				"code": 2002,
				"msg":  err.Error(),
			},
			)
		}

	}
	parse_request = model.Parse_request_struct{
		Pathurl:   pathurl,
		Real_path: real_path,
		Statinfo:  statinfo,
	}
	return parse_request, err
}

func list(c *gin.Context) {
	var request_list model.Parse_request_list
	c.ShouldBind(&request_list)
	parse_request, err := check_pathurl(c)
	if err != nil {
		return
	}
	fb, err := files.IterDirectory(parse_request, request_list.Onlydir)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  err.Error(),
		},
		)
	}
	c.JSON(http.StatusOK, fb)
}

func down(c *gin.Context) {
	parse_request, err := check_pathurl(c)

	if err != nil {
		return
	}
	if parse_request.Statinfo.IsDir() {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  parse_request.Pathurl + " Directory does not support download",
		},
		)
		return
	}
	file_name := parse_request.Statinfo.Name()
	c.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	c.Header("Content-Disposition", "attachment; filename="+file_name)
	//浏览器下载或预览
	c.Header("Content-Disposition", "inline;filename="+file_name)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(parse_request.Real_path)
}

func up(c *gin.Context) {
	parse_request, err := check_pathurl(c)
	if err != nil {
		return
	}
	if !parse_request.Statinfo.IsDir() {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  parse_request.Pathurl + " is not a directory",
		},
		)
		return
	}
	filepart, err := c.FormFile("file")
	filename := filepath.Join(parse_request.Real_path, filepart.Filename)
	err = c.SaveUploadedFile(filepart, filename)

	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  " upload failed",
		})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{
		"code": 200,
		"msg":  "upload ok",
	})
}



func delete(c *gin.Context) {
	parse_request, err := check_pathurl(c)
	if err != nil {
		return
	}
	if files.BeforFileChangeCheck(parse_request.Real_path) {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  "cat not delete {{basedir}}",
		})
		return
	}
	err = os.RemoveAll(parse_request.Real_path)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  files.FileErrPars(err),
		},
		)
		return
	}
	c.AbortWithStatusJSON(200, gin.H{
		"code": 200,
		"msg":  "delete ok",
	},
	)

}

func rename(c *gin.Context) {
	var request_nwename model.Parse_request_rename
	parse_request, err := check_pathurl(c)
	if err != nil {
		return
	}
	if files.BeforFileChangeCheck(parse_request.Real_path) {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  "cat not rename {{basedir}}",
		})
		return
	}
	request_nwename = model.Parse_request_rename{}
	c.ShouldBind(&request_nwename)
	parentDirectory := filepath.Dir(parse_request.Real_path)
	new_realpath := filepath.Join(parentDirectory, request_nwename.Newname)
	err = os.Rename(parse_request.Real_path, new_realpath)
	//zlog.SugLog.Infof("%#v666666666666666\n",err)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  files.FileErrPars(err),
		},
		)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "rename is ok",
	},
	)
}

func mv(c *gin.Context) {
	basedir := config.ViperConfig.GetString("basedir")
	real_path := filepath.Join(basedir, c.GetString("role"))
	var request_mv model.Parse_request_mv
	gin.EnableJsonDecoderDisallowUnknownFields()
	err := c.BindJSON(&request_mv)
	if err != nil {
		c.Error(err)
		c.JSON(200, gin.H{
			"code": 400,
			"err":  "给定的json参数缺少或错误",
		})
		return
	}
	datas := files.Mvproject(real_path, &request_mv)
	datas.Code = 200
	c.JSON(200, datas)
}

func diskinfo(c *gin.Context) {
	basedir := config.ViperConfig.GetString("basedir")
	Diskinfo, e := systeminfo.GetDiskinfo(basedir)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  files.FileErrPars(e),
		})
		return
	}
	c.JSON(200, model.Diskstruct{
		Dmsg: model.Dmsg{Code: 200,
			Msg: "ok"},
		Data: Diskinfo,
	})
}

func createnewdir(c *gin.Context) {
	var request_nwename model.Parse_request_rename
	parse_request, err := check_pathurl(c)
	if err != nil {
		return
	}
	request_nwename = model.Parse_request_rename{}
	c.ShouldBind(&request_nwename)
	new_realpath := filepath.Join(parse_request.Real_path, request_nwename.Newname)
	err = os.Mkdir(new_realpath, os.ModePerm)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"code": 2003,
			"msg":  files.FileErrPars(err),
		},
		)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "create is ok",
	},
	)
}
