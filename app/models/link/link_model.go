// Package link 模型
package link

import (
	"gohub/app/models"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

type Link struct {
	models.BaseModel
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
	models.CommonTimestampsField
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}
