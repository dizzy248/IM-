package models

import (
	"fmt"
	"gin_webSocket_project_IM/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name     string `json:"姓名"`
	PassWord string
	Phone    string `valid:"matches(^1[3-9]{1}\\d{9}$)"` //电话邮箱校验
	Email    string `valid:"email"`
	//身份唯一标识
	Identidy string
	ClientIp string
	//客户端口
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"colum:login_out_time" json:"login_out_time"`
	IsLogout      bool
	//设备信息
	DeviceInfo string
}

func (table *UserBasic) TableName() string {
	return "user_basic" //返回数据库，相当于在数据库中创建一个名叫user_basic的表
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByNameAndPwd(name, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	return user
}

func FindUserByPhone(phone string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone = ?", phone).First(&user)
	return user
}

func FindUserByEmail(email string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("email = ?", email).First(&user)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}
