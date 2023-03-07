package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&LogModel{}, &GoldModel{})
	Db = db
}

type LogStore interface {
	Add(msg string)
	Load(size int) (any, error)
}

// LogModel 日志模型
type LogModel struct {
	gorm.Model
	Message string `json:"message"`
}

// Add 添加日志
func (LogModel) Add(msgs ...any) {
	var msg string
	for _, s := range msgs {
		msg = fmt.Sprintf("%s %v", msg, s)
	}
	Db.Create(&LogModel{Message: msg})
}

// Load 加载日志
func (LogModel) Load(size int) ([]*LogModel, error) {
	var res []*LogModel
	Db.Find(&res).Limit(size)
	return res, nil
}

// GoldModel 日志模型
type GoldModel struct {
	gorm.Model
	Price float64 `json:"price"`
}

// Add 添加数据
func (GoldModel) Add(price float64) {
	Db.Create(&GoldModel{Price: price})
}

// Load 加载数据
func (GoldModel) Load(size int) ([]*GoldModel, error) {
	var res []*GoldModel
	Db.Find(&res).Limit(size)
	return res, nil
}
