package model

import (
	"time"
)

// Treasurehunting
type Treasurehunting struct {
	TreasureID       string    `json:"treasureID" gorm:"primaryKey"`           //寻宝活动的ID
	Poster           string    `gorm:"foreignKey:Student.StuID" json:"poster"` //发布人的ID
	HeadImage        string    `json:"headImage"`                              //发布人的头像
	Text             string    `json:"text"`                                   //寻宝活动的内容
	Treasurelocation int       `json:"treasurelocation"`                       //寻宝活动的地点
	Thing            string    `json:"thing"`                                  //要寻找的物品
	Time             time.Time `json:"time"`                                   //发布时间
	Image            string    `json:"image"`                                  //物品图片
}
