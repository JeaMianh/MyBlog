package model

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {

	// json 标签用于指定 json 编码和解码时的名称
	// gorm 标签用于指定字段在数据库中的映射和约束
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null;column:title;size:255"`
	Author    string         `json:"author" gorm:"not null;column:author;size:100"`
	Content   string         `json:"content" gorm:"not null;column:content;size:255"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Views     uint           `json:"views" gorm:"default:0"`
	Tags      string         `json:"tags" gorm:"type:varchar(100)"`    // 类型为可变长度字符串，最大100
	Category  string         `json:"category" gorm:"type:varchar(50)"` // 和 size 的区别在于是否指定了类型
}
