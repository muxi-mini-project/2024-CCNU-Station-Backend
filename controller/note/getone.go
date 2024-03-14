package note

//
//import (
//	"errors"
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//	"guizizhan/model"
//	response "guizizhan/response/note"
//	"guizizhan/service/tool"
//)
//
//// GetThePostNote 获取特定帖子的接口。
//// @Summary 获取特定帖子
//// @Description 和“获取所有帖子”基本相同，唯一不同的是返回的帖子信息不是数组，而是数组中的一个元素。
//// @ID get-the-post
//// @Produce json
//// @Param postid query string true "帖子的ID"
//// @Security Bearer
//// @Api(tags="获取")
//// @Success 200 {object} response.GetNote_resp
//// @Failure 200 {object} response.GetNote_resp
//// @Router /api/getactivity/thepostnote [get]
//func GetThePostNote(c *gin.Context, db *gorm.DB) {
//	var msg string
//	_, yn := tool.GetStudentID(c)
//	postid := c.Query("postid")
//	var post model.Post
//	var posimage model.PostImage
//	res := db.Model(&model.Post{}).Where(&model.Post{PostID: postid}).First(&post)
//	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
//		msg = "没有这个帖子"
//	} else {
//		msg = "成功获取到帖子"
//	}
//	db.Model(&model.PostImage{}).Where(&model.PostImage{PostID: postid}).First(&posimage)
//	if yn {
//		response.GetTheNote_ok(c, post, posimage, msg)
//	} else {
//		response.GetTheNote_fail(c)
//	}
//
//}
