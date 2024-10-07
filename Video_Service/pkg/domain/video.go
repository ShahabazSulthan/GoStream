package domain

type Video struct {
	ID      uint   `json:"id" gorm:"primarykey;auto_increment"`
	VideoId string `json:"videoid" gorm:"uniqueIndex;not null"`
}
