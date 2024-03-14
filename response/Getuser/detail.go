package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
	"time"
)

type User_data struct {
	Stuid           string    `json:"stuid"`
	RealName        string    `json:"realName"`
	Nickname        string    `json:"nickname"`
	Grade           string    `json:"grade"`
	College         string    `json:"college"`
	Gender          string    `json:"gender"`
	Headimage       string    `json:"headimage"`
	Age             int       `json:"age"`
	Sign            string    `json:"sign"`
	FriendsNumber   int       `json:"friends_number"`
	FollowersNumber int       `json:"fan_number"`
	FollowerNumber  int       `json:"follower_number"`
	PostNumber      int       `json:"post_number"`
	Date            time.Time `json:"date"`
	StayDate        int       `json:"stay_date"`
	YNLogin         bool      `json:"YNLogin"`
	YNSelf          bool      `json:"YNSelf"`
	MBTI            string    `json:"mbti"`
}
type User_resp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data User_data `json:"data"`
}

func GetUserdetails_ok(c *gin.Context, stuid string, realname string, nickname string, grade string, college string, gender string, headimage string, age int, sign string, friendsNumber int, FollowersNumber int, FollowerNumber int, PostNumber int, Date time.Time, StayDate int, YNSelf bool, mbti string) {
	var data = User_data{
		Stuid:           stuid,
		RealName:        realname,
		Nickname:        nickname,
		Grade:           grade,
		College:         college,
		Gender:          gender,
		Headimage:       headimage,
		Age:             age,
		Sign:            sign,
		FriendsNumber:   friendsNumber,
		FollowersNumber: FollowersNumber,
		FollowerNumber:  FollowerNumber,
		PostNumber:      PostNumber,
		Date:            Date,
		StayDate:        StayDate,
		YNSelf:          YNSelf,
		MBTI:            mbti,
		YNLogin:         true,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  "已成功获取",
		"data": data,
	})
}
func GetUserdetails_fail(c *gin.Context) {
	var data = User_data{
		YNLogin: false,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
