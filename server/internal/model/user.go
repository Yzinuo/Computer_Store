package model

import (
	"math/rand"
	"strconv"
	"time"
	"computer_store/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID           int      `gorm:"primaryKey;autoIncrement;column:id;comment:用户唯一标识" json:"id"`
	Username     string    `gorm:"unique;not null;column:username;comment:用户名" json:"nickname"`
	Password     string    `gorm:"not null;column:password;comment:密码（加密存储）"`
	Email        string    `gorm:"unique;not null;column:email;comment:邮箱" json:"email"`
	Introduction string    `gorm:"column:introduction;comment:简介" json:"intro"`
	Avatar       string    `gorm:"column:avatar;comment:头像" json:"avatar"`
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:created_at;comment:创建时间"`
	UpdatedAt    time.Time `gorm:"column:updated_at;comment:更新时间"`
}

func GetUserInfoByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	result := db.First(&user,"email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserInfoById(db *gorm.DB, id int) (*User, error) {
	var user User
	result := db.First(&user,"id =?", id)
	if result.Error!= nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateNewUser(db *gorm.DB, email,password string) error {
	randomNumber := rand.Intn(100)

	// 将随机数转换为字符串
	randomNumberStr := strconv.Itoa(randomNumber)

	pass,_ := utils.BcryptHash(password)
	userinfo := &User{
		Email: email,
		Username:"游客"+randomNumberStr,
		Password: pass,
		Avatar: "https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png",
		Introduction: "万事顺遂",
	}

	result := db.Create(&userinfo)
	if result.Error!= nil {
		return result.Error
	}

	return nil
}

func UpdateUserInfo(db *gorm.DB,id int,username,email,avatar,intro string) (*User,error) {
	userinfo := &User{
		ID: id,
		Username: username,
		Email: email,
		Avatar: avatar,
		Introduction: intro,
	}
	result := db.Model(&User{}).Where("id = ?", id).Updates(userinfo)
	return userinfo,result.Error
}