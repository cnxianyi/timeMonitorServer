package types

import (
	"time"
)

type UploadForm struct {
	Title    string    `json:"title" binding:"required"`
	Process  string    `json:"process" binding:"required"`
	Time     time.Time `json:"time" binding:"required"`
	UserName string    `json:"user_name" binding:"required"`
}

type UploadTimeForm struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	DailyTime uint   `json:"dailyTime" binding:"required"`
	EveryTime uint   `json:"everyTime" binding:"required"`
}

type TitleClassModel struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Model      uint      `json:"model"`   // 1 全等 2 内含
	Content    string    `json:"title"`   // 匹配内容
	Legend     uint      `json:"legend"`  // 返回值 1 学习 2 娱乐 3 社交 4 其他
	Process    uint      `json:"process"` // 1 匹配 2 不匹配
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime"`
}

func (TitleClassModel) TableName() string {
	return "title_class" // Explicitly setting table name
}

type UserModel struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	DailyTime  uint      `json:"daily_time" binding:"required"`
	EveryTime  uint      `json:"every_time" binding:"required"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime"`
}

func (UserModel) TableName() string {
	return "users" // Explicitly setting table name
}

type ProcessModel struct {
	Id         uint         `json:"id" gorm:"primaryKey"`
	Process    string       `json:"process" gorm:"not null;type:varchar(255)"`
	Date       string       `json:"date" gorm:"type:date"`
	Hour       uint8        `json:"hour" gorm:"type:tinyint"`
	UserId     uint         `json:"user_id" gorm:"not null;index"` // Add index for foreign key
	User       UserModel    `gorm:"foreignKey:UserId"`
	CreateTime time.Time    `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time    `json:"update_time" gorm:"autoUpdateTime"`
	Titles     []TitleModel `gorm:"foreignKey:ProcessId"` // ProcessId 是 Title 表中的外键
}

func (ProcessModel) TableName() string {
	return "processes" // Explicitly setting table name
}

type TitleModel struct {
	Id         uint         `json:"id" gorm:"primaryKey"`
	ProcessId  uint         `json:"process_id" gorm:"not null;index"` // Add index for foreign key
	Process    ProcessModel `gorm:"foreignKey:ProcessId"`             // Defines the relationship
	Title      string       `json:"title" gorm:"not null;type:varchar(255)"`
	Time       uint         `json:"time" gorm:"type:int"` // Assuming 'Time' is a duration in minutes/seconds
	CreateTime time.Time    `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time    `json:"update_time" gorm:"autoUpdateTime"`
}

// TableName specifies the table name for TitleModel
func (TitleModel) TableName() string {
	return "titles" // Explicitly setting table name
}

type ProcessResponse struct {
	Process string          `json:"process"`
	Hour    uint8           `json:"hour"`
	Titles  []TitleResponse `json:"titles"`
}

type UploadResponse struct {
	Lave   int    `json:"lave"`
	Notice string `json:"notice"`
}
type TitleResponse struct {
	Title  string `json:"title"`
	Time   uint   `json:"time"`
	Legend uint   `json:"legend"`
}
