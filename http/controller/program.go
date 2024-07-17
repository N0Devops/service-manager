package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"service-manager/http/common"
	"service-manager/http/middleware"
	"service-manager/program"
)

type ProgramQuery struct {
	Name   string `form:"name" json:"name,omitempty"`
	Config string `form:"conf" json:"ci,omitempty"`
}

type ProgramBody struct {
	Data string `form:"data" json:"data,omitempty"`
}

type ProgramController struct {
}

func (controller ProgramController) programs(ctx *gin.Context) interface{} {
	conf := program.List()
	return conf
}

func (controller ProgramController) getProgramAction(ctx *gin.Context) (program.ProgramAction, ProgramQuery) {
	var query ProgramQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		panic(err)
	}
	conf := program.Load()
	p, ok := conf[query.Name]
	if !ok {
		panic(fmt.Errorf("program define not found"))
	}
	return program.NewProgramAction(p), query
}

func (controller ProgramController) start(ctx *gin.Context) interface{} {
	pa, _ := controller.getProgramAction(ctx)
	bytes, err := pa.Start()
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (controller ProgramController) stop(ctx *gin.Context) interface{} {
	pa, _ := controller.getProgramAction(ctx)
	bytes, err := pa.Stop()
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (controller ProgramController) restart(ctx *gin.Context) interface{} {
	pa, _ := controller.getProgramAction(ctx)
	bytes, err := pa.Restart()
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (controller ProgramController) status(ctx *gin.Context) interface{} {
	pa, _ := controller.getProgramAction(ctx)
	bytes, err := pa.Status()
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (controller ProgramController) readConfig(ctx *gin.Context) interface{} {
	pa, query := controller.getProgramAction(ctx)
	bytes, err := pa.ReadConfig(query.Config)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (controller ProgramController) writeConfig(ctx *gin.Context) interface{} {
	pa, query := controller.getProgramAction(ctx)
	var b ProgramBody
	if err := ctx.ShouldBind(&b); err != nil {
		panic(err)
	}
	log.Println(query.Config, b.Data)
	err := pa.WriteConfig(query.Config, []byte(b.Data))
	if err != nil {
		panic(err)
	}
	return nil
}

func (controller ProgramController) Router(r *gin.RouterGroup) {
	res := common.Response{}
	g := r.Group("/program")
	g.Use(middleware.AuthorizationMiddleware())
	g.GET("", res.SafetyWithData(controller.programs))
	g.GET("/start", res.SafetyWithData(controller.start))
	g.GET("/stop", res.SafetyWithData(controller.stop))
	g.GET("/restart", res.SafetyWithData(controller.restart))
	g.GET("/status", res.SafetyWithData(controller.status))
	g.GET("/config", res.SafetyWithData(controller.readConfig))
	g.POST("/config", res.SafetyWithData(controller.writeConfig))
}

func NewProgramController() ProgramController {
	return ProgramController{}
}
