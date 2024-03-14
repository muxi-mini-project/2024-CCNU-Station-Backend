package response

import "github.com/gin-gonic/gin"

type Follow_data struct {
	YNLogin bool `json:"YNLogin"`
}
type Follow_resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data Follow_data `json:"data"`
}

func Follow_ok(c *gin.Context, msg string) {
	var data = Follow_data{
		YNLogin: true,
	}
	c.JSON(200, gin.H{
		"code": SUCCESS,
		"msg":  msg,
		"data": data,
	})
}
func Follow_fail(c *gin.Context, msg string, ynlogin bool) {
	var data = Follow_data{
		YNLogin: ynlogin,
	}
	var msg1 string
	if ynlogin {
		msg1 = ""
	} else {
		msg1 = " 未登录"
	}
	c.JSON(200, gin.H{
		"code": FAIL,
		"msg":  msg + msg1,
		"data": data,
	})
}
