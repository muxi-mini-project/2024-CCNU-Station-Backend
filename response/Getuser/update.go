package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
)

type Update_data struct {
	YNLogin bool `json:"YNLogin"`
}
type Update_resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data Update_data `json:"data"`
}

func Update_ok(c *gin.Context, msg string) {
	var data = Update_data{YNLogin: true}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  msg,
		"data": data,
	})
}
func Update_fail(c *gin.Context, yn bool, msg string) {
	var data = Update_data{
		YNLogin: yn,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  msg,
		"data": data,
	})
}
