package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, res interface{}) {
	// TODO: Logger
	c.JSON(http.StatusOK, res)
}

func ResponseFail(c *gin.Context, errCode int, res interface{}) {
	// TODO: Logger
	c.JSON(errCode, res)
}
