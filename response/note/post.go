package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
)

type Postnote_data struct {
	YNLogin bool `json:"YNLogin"`
}
type Postnote_resp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data Postnote_data `json:"data"`
}

func Postnote_ok(c *gin.Context) {
	var data = Postnote_data{
		YNLogin: true,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  "发帖成功",
		"data": data,
	})
}
func Postnote_fail(c *gin.Context) {
	var data = Postnote_data{
		YNLogin: false,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
