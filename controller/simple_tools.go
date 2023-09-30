package controller

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ToUpperHandler(c *gin.Context) {
	var str = c.Param("str")
	c.String(200, strings.ToUpper(str))
}

func ToLowerHandler(c *gin.Context) {
	var str = c.Param("str")
	c.String(200, strings.ToLower(str))
}
