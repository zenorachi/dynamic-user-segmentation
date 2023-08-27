package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/logger"
)

func newResponse(c *gin.Context, statusCode int, msg any) {
	logger.Info(c.Request.RequestURI, "successfully")
	c.JSON(statusCode, msg)
}

type errorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(c *gin.Context, statusCode int, err string) {
	logger.Error(c.Request.RequestURI, err)
	c.AbortWithStatusJSON(statusCode, errorResponse{Error: err})
}
