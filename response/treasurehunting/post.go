package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
)

type Post_treasurehunting_data struct {
	YNLogin bool `json:"YNLogin"`
}
type Post_treasurehunting_resp struct {
	Code int                       `json:"code"`
	Msg  string                    `json:"msg"`
	Data Post_treasurehunting_data `json:"data"`
}

func Post_treasurehunting_ok(c *gin.Context) {
	var data = Post_treasurehunting_data{
		YNLogin: true,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  "发布寻宝活动成功",
		"data": data,
	})
}
func Post_treasurehunting_fail(c *gin.Context) {
	var data = Post_treasurehunting_data{
		YNLogin: false,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
