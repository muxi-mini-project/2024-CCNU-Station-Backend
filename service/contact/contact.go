package contact

import (
	"errors"
	"gorm.io/gorm"
	"guizizhan/model"
)

func CheckOtherIfFollowYou(ownerID, targetID string, db *gorm.DB) int {
	var c model.Contact
	res := db.Model(&model.Contact{}).Where(&model.Contact{OwnerID: targetID, TargetID: ownerID}).First(&c)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return 0 //0代表对方并没有关注你，你可以关注他
	} else {
		if c.Type == 1 {
			return 1 //1代表对方已经关注你了，并且你们是好友，不能再关注了
		} else if c.Type == 2 {
			return 2 //2代表对方已经关注你了，他是你的粉丝，你可以关注他
		} else {
			return 3 //3代表你是他的粉丝，你不能关注他
		}
	}
}
