package treasurehunting

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model"
	"guizizhan/pkg/qiniu"
	response "guizizhan/response/treasurehunting"
	"guizizhan/service/generateID"
	"guizizhan/service/tool"
	"strconv"
	"time"
)

// PostTreasureHunting 处理发布寻宝活动的请求。
// @Summary 发布寻宝活动
// @Description 返回的信息比较简单，code还是1000表示成功，1000表示失败（未登录），YNLogin代表是否登录，不过code信息已经说明了。
// @ID post-treasure-hunting
// @Accept json
// @Produce json
// @Param image query string true "图片文件的Key"
// @Param content formData string true "活动内容"
// @Param treasurelocation query string true "寻宝地点"
// @Param thing formData string true "寻找的物品"
// @Security Bearer
// @Api(tags="发布")
// @Success 200 {object} response.Post_treasurehunting_resp
// @Failure 200 {object} response.Post_treasurehunting_resp
// @Router /api/post/post_treasure_hunting [post]
func PostTreasureHunting(c *gin.Context, db *gorm.DB) {
	key, _ := c.GetQuery("image")
	URL := qiniu.GenerateURL(key)
	content := c.PostForm("content")
	wherestring, _ := c.GetQuery("where")
	whereint, _ := strconv.Atoi(wherestring)
	thing := c.PostForm("thing")
	stuid, yn := tool.GetStudentID(c)
	student, _ := model.FindStudfromID(stuid, db)
	treasureid := generateID.GenerateTreasureID(db)
	var treasurehunting model.Treasurehunting
	treasurehunting = model.Treasurehunting{
		TreasureID:       treasureid,
		Image:            URL,
		Thing:            thing,
		Treasurelocation: whereint,
		Poster:           stuid,
		HeadImage:        student.HeadImage,
		Time:             time.Now(),
		Text:             content,
	}

	db.Create(&treasurehunting)

	//响应
	if yn {
		response.Post_treasurehunting_ok(c)
	} else {
		response.Post_treasurehunting_fail(c)
	}
}
