package model

type PostImage struct {
	PostID string `json:"postID" gorm:"primaryKey"`
	Image1 string `json:"image1"`
	Image2 string `json:"image2"`
	Image3 string `json:"image3"`
	Image4 string `json:"image4"`
	Image5 string `json:"image5"`
}
