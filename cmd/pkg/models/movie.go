package models

import (
	"github.com/ManivelRamu/CMS-Module/cmd/pkg/config"
	"github.com/ManivelRamu/CMS-Module/cmd/pkg/paginations"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
)

type Claims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

var db *gorm.DB

type Movie struct {
	gorm.Model
	Name     string `json:"name"  gorm:"unique;not null"`
	Director string `json:"director"`
	Rating   string `json:"rating"`
	Year     string `json:"year"`
	Poster   string `json:"poster"`
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

func GetAllMovies(pagination paginations.Pagination) (*paginations.Pagination, error) {
	Movies := []Movie{}
	db.Scopes(config.Paginate(Movies, &pagination, db)).Find(&Movies)
	pagination.Rows = Movies
	return &pagination, nil
}

// func GetAllMovies(pagination paginations.Pagination) (*paginations.Pagination, error) {
// 	movies := []model.Movies{}
// 	conn.Scopes(repo.Paginate(movies, &pagination, conn)).Find(&movies)
// 	pagination.Rows = movies
// 	return &pagination, nil
// }

// type search struct {
// 	db *gorm.DB
// }

// func (productModel search) StartsWith(keyword string) ([]Movie, error) {
// 	rows, err := productModel.db.QueryExpr("select * from product where name like ?", keyword + "%")
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		products := []entities.Movie{}
// 		for rows.Next() {
// 			var id int64
// 			var name string
// 			var price float32
// 			var quantity int
// 			var status bool
// 			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
// 			if err2 != nil {
// 				return nil, err2
// 			} else {
// 				product := entities.Product{id, name, price, quantity, status}
// 				products = append(products, product)
// 			}
// 		}
// 		return products, nil
// 	}
// }

//search
func SearchMovie(arg string, value string) *Movie {
	var search Movie
	db.Where("? LIKE ?", arg, value).Find(&search)
	return &search

}

//search

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

//shows

type Show struct {
	gorm.Model
	ShowName     string `gorm:"foreignkey:ShowName;association_foreignkey:ShowName;unique" json:"showname"`
	ShowDirector string `json:"showdirector"`
	ShowRating   string `json:"showrating"`
	ShowYear     string `json:"showyear"`
	Episodes     string `json:"episodes"`

	//Posters   string `json:"poster"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Show{})
}
func (s *Show) Createshow() *Show {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllshow(pagination paginations.Pagination) (*paginations.Pagination, error) {
	Show := []Show{}
	db.Scopes(config.Paginate(Show, &pagination, db)).Find(&Show)
	pagination.Rows = Show
	return &pagination, nil
}

func GetshowById(Id int64) (*Show, *gorm.DB) {
	var Getshow Show
	db := db.Where("ID=?", Id).Find(&Getshow)
	return &Getshow, db
}

func Deleteshow(ID int64) Show {
	var show Show
	db.Where("ID=?", ID).Delete(show)
	return show
}

//Episode

type Episode struct {
	gorm.Model
	EpisodeNo   string `gorm:"not null; unique" json:"episodeno"`
	EpisodeName string `gorm:"not null; unique" json:"episodename"`
	Description string `json:"description"`
	SName       string `json:"SName"`
	Platform    string `json:"platform"`

	//Posters   string `json:"poster"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Episode{})
}
func (e *Episode) Createepisode() *Episode {
	db.NewRecord(e)
	db.Create(&e)
	return e
}

func GetAllepisode(pagination paginations.Pagination) (*paginations.Pagination, error) {
	episode := []Show{}
	db.Scopes(config.Paginate(episode, &pagination, db)).Find(&episode)
	pagination.Rows = episode
	return &pagination, nil
}

func GetepisodeById(Id int64) (*Episode, *gorm.DB) {
	var Getepisode Episode
	db := db.Where("ID=?", Id).Find(&Getepisode)
	return &Getepisode, db
}

func Deleteepisode(ID int64) Episode {
	var episode Episode
	db.Where("ID=?", ID).Delete(episode)
	return episode
}

//seasons---------------------------------------------------------------
type Season struct {
	gorm.Model

	SeasonNo    string `json:"seasonno"`
	SeasonName  string `gorm:"primary_key"`
	ShowName    string `gorm:"foreignkey:showname;association_foreignkey:sname"`
	Episodes    string `json:"episodes"`
	AvailableOn string `json:"availableon"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Season{})
}
func (s *Season) Createseason() *Season {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllseason(pagination paginations.Pagination) (*paginations.Pagination, error) {
	season := []Season{}
	db.Scopes(config.Paginate(season, &pagination, db)).Find(&season)
	pagination.Rows = season
	return &pagination, nil
}

func GetseasonById(Id int64) (*Season, *gorm.DB) {
	var Getseason Season
	db := db.Where("ID=?", Id).Find(&Getseason)
	return &Getseason, db
}

func Deleteseason(ID int64) Season {
	var season Season
	db.Where("ID=?", ID).Delete(season)
	return season
}

//seasons

//pagination
// type Pagination struct {
// 	Limit int    `json:"limit"`
// 	Page  int    `json:"page"`
// 	Sort  string `json:"sort"`
// }

// func (b *Movie) TableName() string {
// 	return "movies"
// }
