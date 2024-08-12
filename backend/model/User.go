package model

import (
	"backend/weberr"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 添加用户
func CreateUser(data *User) int {
	if db == nil {
		logrus.Errorln("Database instance is nil")
		return weberr.INTERNAL_SERVER_ERROR
	}
	err := db.Create(&data).Error
	if err != nil {
		return weberr.INTERNAL_SERVER_ERROR
	}
	return weberr.SUCCESS
}

// 查询用户是否存在
func CheckUserExists(username string) int {
	var count int64
	if db == nil {
		logrus.Errorln("Database instance is nil")
		return weberr.INTERNAL_SERVER_ERROR
	}
	db.Model(&User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return weberr.USER_ALREADY_EXISTS // 用户已存在
	}
	return weberr.USER_NOT_FOUND // 用户不存在
}

// DeleteUser 删除用户
func DeleteUser(name string) int {
	if db == nil {
		logrus.Error("Database instance is nil")
		return weberr.INTERNAL_SERVER_ERROR
	}
	err := db.Where("username = ? ", name).Unscoped().Delete(&User{}).Error
	//err := db.Delete(&User{}, id).Error
	if err != nil {
		//logrus.Errorf("Failed to delete user with ID %d: %v", name, err)
		return weberr.INTERNAL_SERVER_ERROR
	}

	return weberr.SUCCESS
}
