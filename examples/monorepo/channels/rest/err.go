package rest

import (
	"{{.moduleName}}/pkg/domains/errs"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// HandleError handles errors provided by the application
// and returns the appropriate response to the client
// based on the error type
func HandleError(
	ctx *gin.Context,
	err error,
) {
	if err == nil {
		zap.L().Warn("error is nil")
		ctx.JSON(500, gin.H{
			"error": "internal server error",
		})
		return
	}

	obj, ok := err.(*errs.Entity)
	if !ok {
		zap.L().Error("error is not of type Entity", zap.Error(err))
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(obj.Code, obj)
}
