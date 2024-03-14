package user_behave

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model"
	"guizizhan/pkg/CCNU"
	"guizizhan/pkg/token"
	"guizizhan/response"
)

// Login 处理用户登录请求的控制器。
// @Summary 用户登录
// @Description 登录失败的时候code为1001（FAILURE），这时first（判断是否第一次登录）会默认false，而token（令牌）会为空，登录成功的话，first反映是否第一次登录,而token则是令牌，你需要把他加到后续的请求头上。
// @ID user-login
// @Accept json
// @Produce json
// @Api(tags="登录")
// @Param stuid formData string true "学号"
// @Param password formData string true "密码"
// @Success 200 {object} response.Loginresp
// @Failure 200 {object} response.Loginresp
// @Router /api/login [post]
func Login(c *gin.Context, db *gorm.DB) {
	stuid := c.PostForm("stuid")
	password := c.PostForm("password")
	fmt.Println(stuid, password)

	// 从数据库中查找学生信息
	s, ok := model.FindStudfromID(stuid, db)

	// 如果学生信息不存在且登录成功（即第一次登录）
	if ok == false && CCNU.LoginSuccess(stuid, password) == true {
		//fmt.Println(stuid, "登录成功且是第一次登录")
		s.AddStudent(stuid, password, db)            // 添加新用户
		tokenString, _ := token.GenerateToken(stuid) // 产生一个 TOKEN
		fmt.Println("tokenString:", tokenString)
		response.Login_ok(c, tokenString, true)
	} else if ok == true && CCNU.LoginSuccess(stuid, password) == true {
		// 如果学生信息存在且登录成功
		//fmt.Println(stuid, "登录成功")
		tokenString, _ := token.GenerateToken(s.StuID) // 产生一个 TOKEN
		fmt.Println("tokenString:", tokenString)

		//更新入学天数
		model.UpdateDate(db, stuid)

		response.Login_ok(c, tokenString, false)
	} else {
		// 登录失败的情况
		response.Login_fail(c)
	}
}
