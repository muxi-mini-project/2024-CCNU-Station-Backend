package tool

import (
	"github.com/gin-gonic/gin"
)

func GetStudentID(c *gin.Context) (string, bool) {
	stuidany, yn := c.Get("stuid")
	stuid := stuidany.(string)
	return stuid, yn
}
