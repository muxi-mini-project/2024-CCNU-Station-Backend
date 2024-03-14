package user_behave

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model"
	"guizizhan/pkg/qiniu"
	response "guizizhan/response/Getuser"
	"guizizhan/service/tool"
	"strconv"
	"time"
)

// GetUserDetails 获取用户详细信息的控制器。
// @Summary 获取用户信息
// @Description 有几个我要解释一下，YNSelf反映的是该用户是否是正在登录的用户，fan_number是粉丝数，follower_number是关注数，nickname是昵称，post_number是发帖数，sign是签名，college是学院
// @ID get-user-details
// @Produce json
// @Param userid query string true "用户ID"
// @Security Bearer
// @Api(tags="获取")
// @Success 200 {object} response.User_resp
// @failure 200 {object} response.User_resp
// @Router /api/user/detail [get]
func GetUserDetails(c *gin.Context, db *gorm.DB) {
	var student model.Student
	// 从URL中获取该用户的ID
	userid, _ := c.GetQuery("userid")
	fmt.Println("userid", userid)
	stuid, yn := c.Get("stuid") // 从TOKEN中获取登录者的ID
	fmt.Println("stuid", stuid)
	student, ok := model.FindStudfromID(userid, db)
	if !ok {
		err := errors.New("未找到这个ID")
		fmt.Println(err)
	}
	if yn {
		response.GetUserdetails_ok(c, student.StuID, student.RealName, student.Nickname, student.Grade, student.College, student.Gender, student.HeadImage, student.Age, student.Sign, student.FriendsNumber, student.FanNumber, student.FollowerNumber, student.PostNumber, student.SchoolDate, student.StayDate, stuid == userid, student.MBTI)
	} else {
		response.GetUserdetails_fail(c)
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"stuid":            student.StuID,
	//	"password":         student.Password,
	//	"realName":         student.RealName,
	//	"nickname":         student.Nickname,
	//	"grade":            student.Grade,
	//	"college":          student.College,
	//	"gender":           student.Gender,
	//	"headimage":        student.HeadImage,
	//	"age":              student.Age,
	//	"sign":             student.Sign,
	//	"friends_number":   student.FriendsNumber,
	//	"followers_number": student.FanNumber,
	//	"follower_number":  student.FollowerNumber,
	//	"post_number":      student.PostNumber,
	//	"date":             student.SchoolDate,
	//	"stay_date":        student.StayDate,
	//	"YNLogin":          yn,              // 如果yn==true,表示账号已登录
	//	"YNSelf":           stuid == userid, // 如果userid与stuid相等就表示这个user就是本人
	//})

}

// UpdateUserInfo 更新用户信息的控制器。
// @Summary 更新用户信息
// @Description 处理更新用户信息的请求，包括更新昵称、年龄、入学日期、个性签名等信息。
// @ID update-user-info
// @Accept json
// @Produce json
// @Param nickname formData string true "昵称"
// @Param age formData string true "年龄"
// @Param date formData string true "入学日期，格式为YYYY-MM-DD"
// @Param sign formData string true "个性签名"
// @Param mbti formData string true "mbti"
// @Security Bearer
// @Api(tags="更新")
// @Success 200 {object} response.Update_resp
// @Failure 200 {object} response.Update_resp
// @Router /api/user/update [post]
func UpdateUserInfo(c *gin.Context, db *gorm.DB) {
	var msg string
	stuid, yn := c.Get("stuid")
	student, _ := model.FindStudfromID(stuid.(string), db)
	nickname := c.PostForm("nickname")
	sign := c.PostForm("sign")
	mbti := c.PostForm("mbti")
	// Convert age to int
	ageStr := c.PostForm("age")
	age, err := strconv.Atoi(ageStr)

	if err != nil {
		msg = "更新失败"
		response.Update_fail(c, yn, msg)
		return
	}

	// Convert date to time.Time
	dateStr := c.PostForm("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		msg = "更新失败"
		response.Update_fail(c, yn, msg)
		return
	}
	today := time.Now()
	// 计算天数差
	daysDifference := int(date.Sub(today).Hours() / 24)

	student.Updateinformation(nickname, sign, mbti, age, daysDifference, date, db)

	msg = "更新成功"

	if yn {
		response.Update_ok(c, msg)
	}

}

// UpdateUserAvatar 更新用户头像的控制器。
// @Summary 更新用户头像
// @Description 你可以由前面的推出，嘿嘿，偷懒，见谅。。。
// @ID update-user-avatar
// @Accept json
// @Produce json
// @Param image query string true "头像文件的Key"
// @Security Bearer
// @Api(tags="更新")
// @Success 200 {object} response.Update_resp
// @Failure 200 {object} response.Update_resp
// @Router /api/user/avatar [get]
func UpdateUserAvatar(c *gin.Context, db *gorm.DB) {
	var msg string

	stuid, yn := tool.GetStudentID(c)

	if stuid == "" {
		msg = "stuid is empty"
		response.Update_fail(c, yn, msg)
		return
	}

	_, ok := model.FindStudfromID(stuid, db)
	if !ok {
		msg = "没有找到这个用户"
		response.Update_fail(c, yn, msg)
	}
	key, _ := c.GetQuery("image")

	url := qiniu.GenerateURL(key)

	model.Updateheadimage(db, stuid, url)

	msg = "保存成功"
	response.Update_ok(c, msg)
}
