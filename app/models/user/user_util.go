package user

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

// 判断email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
