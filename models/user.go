package models

import (
	"database/sql"
	"gin/lib"
	"time"
)

func init() {
	lib.Gorm.AutoMigrate(&User{})
}

type User struct {
	ID           uint `gorm:"primary_key"`
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time `gorm:"type:date"`
	Email        string     `gorm:"type:varchar(100);unique_index"`
	Role         string
	MemberNumber string `gorm:"type:varchar(20);unique;not null"` // 设置会员号（member number）唯一并且不为空
	Address      string `gorm:"type:text"`                        // 给address字段创建名为addr的索引
	IgnoreMe     int    `gorm:"-"`                                // 忽略本字段
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}
