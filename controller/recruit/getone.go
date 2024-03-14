package recruit

//
//import (
//	"errors"
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//	"guizizhan/model"
//	response "guizizhan/response/recruit"
//	"guizizhan/service/tool"
//)
//
//// GetTheRecruit 获取特定招募活动的接口。
//// @Summary 获取特定招募活动
//// @Description 和“获取所有招募活动”基本相同，唯一不同的是返回的招募活动信息不是数组，而是数组中的一个元素。
//// @ID get-the-recruit
//// @Produce json
//// @Param recruitID query string true "招募活动的ID"
//// @Security Bearer
//// @Api(tags="获取")
//// @Success 200 {object} response.GetRecruit_resp
//// @Failure 200 {object} response.GetRecruit_resp
//// @Router /api/getactivity/therecruit [get]
//func GetTheRecruit(c *gin.Context, db *gorm.DB) {
//	var msg string
//	_, yn := tool.GetStudentID(c)
//	recruitID := c.Query("recruitID")
//	var recruit model.Recruit
//	res := db.Model(&model.Recruit{}).Where(&model.Recruit{RecruitID: recruitID}).First(&recruit)
//	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
//		msg = "没有这个招募活动"
//	} else {
//		msg = "成功获取"
//	}
//
//	if yn {
//		response.GetTheRecruit_ok(c, recruit, msg)
//	} else {
//		response.GetTheRecruit_fail(c)
//	}
//
//}
