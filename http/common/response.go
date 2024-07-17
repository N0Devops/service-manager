package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct{}

func (r Response) onError(ctx *gin.Context) {
	if rec := recover(); rec != nil {
		switch rec.(type) {
		case CodeError:
			err := rec.(CodeError)
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code": err.Code,
				"err":  err.Err.Error(),
			})
		case error:
			err := rec.(error)
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code": -1,
				"err":  err.Error(),
			})
		default:
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code": -1,
				"err":  rec,
			})
		}

	}
}

func (r Response) SafetyWithData(callback func(ctx *gin.Context) interface{}) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer r.onError(ctx)
		v := callback(ctx)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"data": v,
		})
	}
}
