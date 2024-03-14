package treasurehunting

//
//import (
//	"errors"
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//	"guizizhan/model"
//	response "guizizhan/response/treasurehunting"
//	"guizizhan/service/tool"
//)
//
//// GetTheTreasureHunting 获取特定寻宝活动的接口。
//// @Summary 获取特定寻宝活动
//// @Description 和“获取所有寻宝活动”基本相同，唯一不同的是返回的寻宝活动信息不是数组，而是数组中的一个元素。
//// @ID get-the-treasure-hunting
//// @Produce json
//// @Param treasureID query string true "寻宝活动的ID"
//// @Security Bearer
//// @Api(tags="获取")
//// @Success 200 {object} response.Treasurehunting_resp
//// @Failure 200 {object} response.Treasurehunting_resp
//// @Router /api/getactivity/thetreasurehunting [get]
//func GetTheTreasureHunting(c *gin.Context, db *gorm.DB) {
//	var msg string
//	_, yn := tool.GetStudentID(c)
//	treasureID := c.Query("treasureID")
//	var TreasureHunting model.Treasurehunting
//	res := db.Model(&model.Treasurehunting{}).Where(&model.Treasurehunting{TreasureID: treasureID}).First(&TreasureHunting)
//
//	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
//		msg = "没有这个寻宝活动"
//	} else {
//		msg = "成功获取"
//	}
//
//	if yn {
//		response.GetTheTreasurehunting_ok(c, TreasureHunting, msg)
//	} else {
//		response.GetTheTreasurehunting_fail(c)
//	}
//}
