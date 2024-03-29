package topic

import (
	"gohub/pkg/database"
	"gorm.io/gorm/clause"
)

func Get(idstr string) (topic Topic) {
	database.DB.Preload(clause.Associations).Where("id", idstr).First(&topic) //附带 user 和 category 关联信息
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", field, value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
