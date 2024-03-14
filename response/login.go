package response

import "github.com/gin-gonic/gin"

const (
	SUCCESS            = 1000 // 请求成功 进入前端处理逻辑
	FAIL               = 1001 // 请求错误 前端会自动抛出异常
	REFRESH_CAPTCHA    = 1002 // 需要前端手动判断code == 1002处理的失败
	AUTHORIZATION_FAIL = 1004 // 鉴权失败 前端会自动抛出异常并退出登录
)

type Logindata struct {
	First bool   `json:"First"`
	Token string `json:"token"`
}
type Loginresp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg:"`
	Data Logindata `json:"data"`
}

func Login_ok(c *gin.Context, token string, first bool) {
	var res = Logindata{
		First: first,
		Token: token,
	}
	var msg string
	if first {
		msg = "登录成功且是第一次登录"
	} else {
		msg = "登陆成功但不是第一次登录"
	}
	c.JSON(200, gin.H{
		"code": SUCCESS,
		"msg":  msg,
		"data": res,
	})
}
func Login_fail(c *gin.Context) {
	var res = Logindata{
		First: false,
		Token: "",
	}
	c.JSON(200, gin.H{
		"code": FAIL,
		"msg":  "登录失败",
		"data": res,
	})
}
