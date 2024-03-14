package model

import (
	"errors"
	"gorm.io/gorm"
)

type Achievement struct {
	OwnerID  string `json:"owner" gorm:"primaryKey"`
	Finished string `json:"finished"`
}

func FindAchievement(ownerID string, db *gorm.DB) (string, error) {
	var finished string
	var err error
	var achi Achievement
	res := db.Model(&Achievement{}).Where(&Achievement{OwnerID: ownerID}).First(&achi)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = errors.New("寻找成就时未发现此人")
		return "", err
	}
	finished = achi.Finished
	return finished, nil
}
func UpdateAchievement(ownerID string, achID int, db *gorm.DB) {
	finished, _ := FindAchievement(ownerID, db)
	var newstring string
	if finished[achID-1] == '1' {
		newstring = finished[:achID-1] + "0" + finished[achID:]
	} else {
		newstring = finished[:achID-1] + "1" + finished[achID:]
	}
	db.Model(&Achievement{}).Where(&Achievement{OwnerID: ownerID}).Updates(&Achievement{Finished: newstring})
}

func GetBeginAchi() string {
	var finished string
	for i := 0; i < 100; i++ {
		finished += "0"
	}
	return finished
}

func CreateAchievement(ID string, db *gorm.DB) {
	//初始化成就
	finished := GetBeginAchi()
	achi := Achievement{
		OwnerID:  ID,
		Finished: finished,
	}

	db.Create(&achi)
}
