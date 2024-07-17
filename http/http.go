package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-manager/assets"
	"service-manager/config"
	"service-manager/http/controller"
	"service-manager/http/middleware"
)

func Boot() {
	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.CORSMiddleware())

	root := e.Group("/api")
	controller.NewProgramController().Router(root)
	controller.NewAccountController().Router(root)

	e.NoRoute(func(ctx *gin.Context) {
		assets.HttpHandler(ctx.Request, ctx.Writer)
	})

	conf := config.Load()
	if err := http.ListenAndServe(conf.Http.Addr, e); err != nil {
		panic(err)
	}
}
