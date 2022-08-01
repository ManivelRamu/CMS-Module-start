package config

import (
	"math"

	"github.com/ManivelRamu/CMS-Module/cmd/pkg/paginations"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:levincruze@12@/movies?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = d
}
func GetDB() *gorm.DB {
	return DB
}

//pagination

func Paginate(value interface{}, pagination *paginations.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
