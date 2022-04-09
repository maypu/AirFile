package model

import (
	"gorm.io/gorm"
	"time"
)

// Model gorm.Model的定义
type Model struct {
}

type Common struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Remarks   string         `gorm:"comment:'备注'"`
	Status    string         `gorm:"default:1;comment:'状态'"`
	IpAddress string         `gorm:"comment:'客户端IP'"`
}

// User 用户
type User struct {
	UUID     string
	Nickname string
	Mail     string
	Password string
	Common
}

// File 文件列表
type File struct {
	User         uint      `gorm:"comment:'上传用户'"`
	FileName     string    `gorm:"comment:'原始文件名称'"`
	Path         string    `gorm:"comment:'文件所在路径'"`
	RandomCode   string    `gorm:"comment:'下载验证码'"`
	Password     string    `gorm:"comment:'下载密码'"`
	NumDownloads int       `gorm:"comment:'已下载次数'"`
	LimitTimes   int       `gorm:"comment:'限制次数'"`
	ExpiryTime   time.Time `gorm:"comment:'过期时间'"`
	Common
}

// FileExt 文件扩展
type FileExt struct {
	File         uint
	Password     string
	NumDownloads int       `gorm:"comment:'已下载次数'"`
	LimitTimes   int       `gorm:"comment:'限制次数'"`
	ExpiryTime   time.Time `gorm:"comment:'过期时间'"`
	Common
}

// Text 文字列表
type Text struct {
	User        uint   `gorm:"comment:'上传用户'"`
	Description string `gorm:"comment:'描述'"`
	Password    string
	Common
}

type TextExt struct {
	Text     uint
	Password string
	Common
}

// CornHistory 任务历史
type CornHistory struct {
	DeleteFile uint `gorm:"comment:'删除的文件'"`
	Common
}