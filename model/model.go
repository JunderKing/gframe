package model

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

type Model struct {
	Id        uint           `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	// 获取数据库配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.db"))
	// 连接数据库
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("database.prefix"),
			SingularTable: true,
		},
	})
	// 表迁移
	Db.AutoMigrate(&Demo{})

	if err != nil {
		panic(err)
	}
}

func Create(model interface{}) {
	Db.Create(model)
}

func FirstByField(model interface{}, field string) {
	Db.Where(model, field).First(model)
}

func FindByField(models []interface{}, model *Model, field string) {
	Db.Where(model, field).Find(models)
}
