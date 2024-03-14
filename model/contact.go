package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerID  string `json:"owner_id"`  //A
	TargetID string `json:"target_id"` //B
	Type     int    `json:"type"`      //1表示B是A的friend,A是B的好友,2表示B是A的关注，3表示B是A的粉丝
	Desc     string `json:"desc"`      //预留字段
}

// 关系升级
func ImproveContact(A, B Student, db *gorm.DB, stuid string, UserID string) {
	db.Create(&Contact{
		OwnerID:  stuid,
		TargetID: UserID,
		Type:     2,
	})
	db.Create(&Contact{
		OwnerID:  UserID,
		TargetID: stuid,
		Type:     3,
	})
	db.Model(&Student{}).Where(&Student{StuID: stuid}).Updates(&Student{FollowerNumber: A.FollowerNumber + 1})
	db.Model(&Student{}).Where(&Student{StuID: UserID}).Updates(&Student{FanNumber: B.FanNumber + 1})
}

// 创建关系
func CreateContact(A, B Student, db *gorm.DB, stuid string, UserID string) {
	db.Model(&Contact{}).Where(&Contact{OwnerID: stuid, TargetID: UserID}).Updates(&Contact{Type: 1})
	db.Model(&Contact{}).Where(&Contact{OwnerID: UserID, TargetID: stuid}).Updates(&Contact{Type: 1})
	db.Model(&Student{}).Where(&Student{StuID: stuid}).Updates(&Student{FollowerNumber: A.FollowerNumber + 1, FriendsNumber: A.FriendsNumber + 1})
	db.Model(&Student{}).Where(&Student{StuID: UserID}).Updates(&Student{FanNumber: B.FanNumber + 1, FriendsNumber: B.FriendsNumber + 1})
}
