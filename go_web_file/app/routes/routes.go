package routes

import (
	"filrserver/app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {

}

func Load(r *gin.Engine) {

	// 资源路径
	//r.Static("resources/assets", "./resources/assets")

	// 无权限路由组
	noAuthRouter := r.Group("/")
	{
		noAuthRouter.Any("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"title": "Main website",
			})
		})
		noAuthRouter.Any("/login", authHandler)
	}

	// 权限路由组
	authRouter := r.Group("/").Use(middlewares.JWTAuth())
	{
		authRouter.GET(("/list/*pathurl"), index)
		authRouter.GET(("/down/*pathurl"), down)
		authRouter.Any(("/up/*pathurl"), middlewares.Check_notMethod("POST"), up)
		authRouter.Any(("/del/*pathurl"), middlewares.Check_notMethod("POST"), delete)
		authRouter.POST(("/reset/*pathurl"), rename)
		authRouter.POST(("/create/*pathurl"), createnewdir)
		authRouter.GET(("/sysinfo/disk"), diskinfo)
	}
}
