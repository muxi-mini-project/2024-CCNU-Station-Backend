package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
)

type Postrecruit_data struct {
	YNLogin bool `json:"YNLogin"`
}
type Postrecruit_resp struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data Postrecruit_data `json:"data"`
}

func Postrecruit_ok(c *gin.Context) {
	var data = Postrecruit_data{
		YNLogin: true,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  "发布招募信息成功",
		"data": data,
	})
}
func Postrecruit_fail(c *gin.Context) {
	var data = Postrecruit_data{
		YNLogin: false,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
