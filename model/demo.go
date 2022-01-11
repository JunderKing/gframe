package model

type Demo struct {
	Model
	Id   int
	Name string `gorm:"type:varchar(100);not null;default:''"`
	Sex  int    `gorm:"type:tinyint unsigned;not null;default:0"`
	Age  int    `gorm:"type:tinyint unsigned;not null;default:0"`
}
