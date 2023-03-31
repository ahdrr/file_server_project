package appcmd

import (
	"filrserver/app/middlewares"
	"filrserver/app/routes"
	"filrserver/pkgs/config"
	"filrserver/pkgs/zlog"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Basedir string

func Start() {
	config.Init()
	zlog.Init()
	listenport := config.ViperConfig.GetString("port")
	initGin := initRouter()
	zlog.SugLog.Infof("******服务初始化完成,监听端口为: %v******", listenport)
	err := initGin.Run("0.0.0.0:" + listenport)
	zlog.Fatalerror(err)
}

func initRouter() *gin.Engine {
	// 检查工作目录
	basedir := config.ViperConfig.GetString("basedir")

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

func checkInitDir(real_path string) {
	f, err := os.Stat(real_path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(real_path, os.ModePerm)
		} else {
			zlog.Fatalerror(err)
		}
	} else {
		if !f.IsDir() {
			zlog.Logc.Fatal("初始化错误", zap.String(real_path, "已经存在同名文件,无法创建目录"))
		}
	}

}

func init_role_dir(basedir string) {
	checkInitDir(basedir)
	listnum := len(config.Users.Users)
	config.Users.UserDirsMap = make(map[string]bool, listnum+1)
	config.Users.UserRoleMap = make(map[string]string, listnum)
	for _, u := range config.Users.Users {
		real_path := filepath.Join(basedir, u.Role)
		config.Users.UserDirsMap[real_path] = true
		config.Users.UserRoleMap[u.Username] = u.Role
		go checkInitDir(real_path)
	}
	config.Users.UserDirsMap[basedir] = true
}
