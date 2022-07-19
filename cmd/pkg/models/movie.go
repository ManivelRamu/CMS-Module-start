package models

import (
	"github.com/ManivelRamu/CMS-Module/cmd/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Name     string `json:"name"`
	Director string `json:"director"`
	Rating   string `json:"rating"`
	Year     string `json:"year"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func (m *Movie) CreateMovie() *Movie {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetAllMovies() []Movie {
	var Movies []Movie
	db.Find(&Movies)
	return Movies
}

func GetMoviesById(Id int64) (*Movie, *gorm.DB) {
	var getMovie Movie
	db := db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovie(ID int64) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}
