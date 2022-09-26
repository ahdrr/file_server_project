package appcmd

import (
	"filrserver/app/middlewares"
	"filrserver/app/routes"
	"filrserver/pkgs/config"
	"filrserver/pkgs/zlog"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var Basedir string

func Start() {
	config.Init()
	zlog.Init()
	listenport := config.ViperConfig.GetString("port")
	initGin := initRouter()
	zlog.SugLog.Infof("服务初始化完成,监听端口为: %v", listenport)
	err := initGin.Run("0.0.0.0:" + listenport)
	zlog.Fatalerror(err)
}

func initRouter() *gin.Engine {
	// 检查工作目录
	basedir := config.ViperConfig.GetString("basedir")

	_, err := os.Stat(basedir)
	zlog.Fatalerror(err)
	//设置运行模式
	var ginmode string = config.ViperConfig.GetString("ginmode")
	if ginmode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if ginmode != "debug" {
		zlog.SugLog.Fatalf("运行级别ginmode设置错误，无法识别%v", ginmode)
	}
	//初始化权限目录
	init_role_dir(basedir)
	// 初始化引擎
	r := gin.New()

	// 公共中间件
	r.Use(middlewares.GinLogger())

	r.Use(middlewares.GinRecovery())

	r.Use(middlewares.CORSMiddleware())
	routes.Load(r)
	return r
}

func init_role_dir(basedir string) {
	for _, u := range config.Users.Users {
		go func(role string) {
			real_path := filepath.Join(basedir, role)
			_, err := os.Stat(real_path)
			if err != nil && os.IsNotExist(err) {
				os.MkdirAll(real_path, os.ModePerm)
			} else {
				zlog.Fatalerror(err)
			}
		}(u.Role)
	}

}
