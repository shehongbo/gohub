package user

import "gohub/app/models"

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"` //指示 JSON 解析器忽略字段
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
