package user_behave

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model"
	"guizizhan/response"
	"guizizhan/service/contact"
	"guizizhan/service/tool"
	"net/http"
)

// Follow 处理关注操作的控制器。
// @Summary 关注操作
// @Description 返回信息很少，前面的基本都涵盖了，唯一有点区别的是，只有成功的情况，当且仅当用户登录且成功关注才能成功（1001），所以失败的时候，你只用说“可能未登录或者关注操作非法”。
// @ID follow-action
// @Produce json
// @Param userid query string true "被关注用户的ID"
// @Security Bearer
// @Api(tags = "关注操作”)
// @Success 200 {object} response.Follow_resp
// @Failure 200 {object} response.Follow_resp
// @Router /api/other/follow [get]
func Follow(c *gin.Context, db *gorm.DB) {
	var msg string
	var YNBecomeFriend bool

	// 首先从URL中获取此人的ID
	UserID := c.Query("userid")
	fmt.Println("following", " userid: ", UserID)
	// 再从token中获取自己的ID
	stuid, yn := tool.GetStudentID(c)

	A, _ := model.FindStudfromID(stuid, db)

	B, _ := model.FindStudfromID(UserID, db)

	if stuid == "" || UserID == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "stuid or userid is empty",
		})
		return
	}
	if stuid == UserID {
		c.JSON(http.StatusOK, gin.H{
			"msg": "stuid == userid",
		})
		return
	}

	msg, YNBecomeFriend = statehandle(A, B, db, stuid, UserID)

	if yn && YNBecomeFriend {
		response.Follow_ok(c, msg)
	} else {
		response.Follow_fail(c, msg, yn)
	}
}
func statehandle(A, B model.Student, db *gorm.DB, stuid, UserID string) (string, bool) {
	var msg string
	var YNBecomeFriend bool
	// 先查看你关注的人有没有关注你
	switch contact.CheckOtherIfFollowYou(stuid, UserID, db) {
	case 0: // 0代表对方并没有关注你，你可以关注他,然后他就成了你的关注，你成了他的粉丝
		msg = "已关注，你已成为他的粉丝，他成为你的关注"
		model.ImproveContact(A, B, db, stuid, UserID)
		YNBecomeFriend = false
	case 1: // 1  代表对方已经关注你了，并且你们是好友，不能再关注了
		msg = "无法关注，你们已经相互关注了"
		YNBecomeFriend = true
	case 2: // 2  代表对方已经关注你了，他是你的粉丝，你可以关注他,关注后，你们的关系将是朋友
		msg = "你们已经成为好友了"
		model.CreateContact(A, B, db, stuid, UserID)
		YNBecomeFriend = true
	case 3: // 3  代表你是他的粉丝，但也只是你是他的粉丝，你不能关注他
		msg = "无法关注，你已经是他的粉丝了"
		YNBecomeFriend = false
	}
	return msg, YNBecomeFriend
}
