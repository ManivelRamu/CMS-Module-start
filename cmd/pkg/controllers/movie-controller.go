package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ManivelRamu/CMS-Module/cmd/pkg/models"
	"github.com/ManivelRamu/CMS-Module/cmd/pkg/paginations"
	"github.com/ManivelRamu/CMS-Module/cmd/pkg/utils"
	"github.com/gorilla/mux"
)

var NewMovie models.Movie

// func GetMovie(w http.ResponseWriter, r *http.Request) {
// 	newMovies := models.GetAllMovies()
// 	res, _ := json.Marshal(newMovies)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func GetMovie(w http.ResponseWriter, r *http.Request) {
	limit, er := strconv.Atoi(r.URL.Query()["limit"][0])
	page, er1 := strconv.Atoi(r.URL.Query()["page"][0])

	if er != nil || er1 != nil {
		fmt.Println("Error")
	}
	var paginate = paginations.Pagination{}
	paginate.Limit = limit
	paginate.Page = page
	movies, err := models.GetAllMovies(paginate)
	if movies == nil || err != nil {
		w.Write([]byte("No movies found or Error Happened"))
	}
	res, _ := json.Marshal(movies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing with id")
	}
	movieDetails, _ := models.GetMoviesById(ID)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//search
func SearchMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attribute := vars["attribute"]
	value := vars["value"]
	att := attribute
	val := value
	// if err!=nil {
	// 	fmt.Println("error while searching")
	// }
	movieDetails := models.SearchMovie(att, val)
	res, err := json.Marshal(movieDetails)
	if err != nil {
		fmt.Println("error while searching")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	b := CreateMovie.CreateMovie()
	res, err := json.Marshal(b)
	if err != nil {
		fmt.Println(" Movie Already exists")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movie := models.DeleteMovie(ID)
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var updateMovie = &models.Movie{}
	utils.ParseBody(r, updateMovie)
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")

	}
	movieDetails, db := models.GetMoviesById(ID)
	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}
	if updateMovie.Director != "" {
		movieDetails.Director = updateMovie.Director
	}
	if updateMovie.Rating != "" {
		movieDetails.Rating = updateMovie.Rating
	}
	if updateMovie.Year != "" {
		movieDetails.Year = updateMovie.Year
	}
	if updateMovie.Poster != "" {
		movieDetails.Poster = updateMovie.Poster
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//show--------------------------------------------------------------

var Newshow models.Show

func Getshow(w http.ResponseWriter, r *http.Request) {
	limit, er := strconv.Atoi(r.URL.Query()["limit"][0])
	page, er1 := strconv.Atoi(r.URL.Query()["page"][0])

	if er != nil || er1 != nil {
		fmt.Println("Error")
	}
	var paginate = paginations.Pagination{}
	paginate.Limit = limit
	paginate.Page = page
	show, err := models.GetAllshow(paginate)
	if show == nil || err != nil {
		w.Write([]byte("No movies found or Error Happened"))
	}
	res, _ := json.Marshal(show)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetshowById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	showId := vars["showId"]
	ID, err := strconv.ParseInt(showId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	showDetails, _ := models.GetshowById(ID)
	res, _ := json.Marshal(showDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Createshow(w http.ResponseWriter, r *http.Request) {
	Createshow := &models.Show{}
	utils.ParseBody(r, Createshow)
	b := Createshow.Createshow()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Deleteshow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	showId := vars["showId"]
	ID, err := strconv.ParseInt(showId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	show := models.Deleteshow(ID)
	res, _ := json.Marshal(show)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateShow(w http.ResponseWriter, r *http.Request) {
	var updateshow = &models.Show{}
	utils.ParseBody(r, updateshow)
	vars := mux.Vars(r)
	showId := vars["showId"]
	ID, err := strconv.ParseInt(showId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")

	}
	showDetails, db := models.GetshowById(ID)
	if updateshow.ShowName != "" {
		showDetails.ShowName = updateshow.ShowName
	}
	if updateshow.ShowDirector != "" {
		showDetails.ShowDirector = updateshow.ShowDirector
	}
	if updateshow.ShowRating != "" {
		showDetails.ShowRating = updateshow.ShowRating
	}
	if updateshow.ShowYear != "" {
		showDetails.ShowYear = updateshow.ShowYear
	}
	if updateshow.Episodes != "" {
		showDetails.Episodes = updateshow.Episodes
	}
	db.Save(&showDetails)
	res, _ := json.Marshal(showDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Episode
var Newepisode models.Episode

func Getepisode(w http.ResponseWriter, r *http.Request) {
	limit, er := strconv.Atoi(r.URL.Query()["limit"][0])
	page, er1 := strconv.Atoi(r.URL.Query()["page"][0])

	if er != nil || er1 != nil {
		fmt.Println("Error")
	}
	var paginate = paginations.Pagination{}
	paginate.Limit = limit
	paginate.Page = page
	episode, err := models.GetAllepisode(paginate)
	if episode == nil || err != nil {
		w.Write([]byte("No movies found or Error Happened"))
	}
	res, _ := json.Marshal(episode)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetepisodeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	ID, err := strconv.ParseInt(episodeId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	episodeDetails, _ := models.GetepisodeById(ID)
	res, _ := json.Marshal(episodeDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Createepisode(w http.ResponseWriter, r *http.Request) {
	Createepisode := &models.Episode{}
	utils.ParseBody(r, Createepisode)
	b := Createepisode.Createepisode()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Deleteepisode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	ID, err := strconv.ParseInt(episodeId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	episode := models.Deleteepisode(ID)
	res, _ := json.Marshal(episode)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Updateepisode(w http.ResponseWriter, r *http.Request) {
	var updateepisode = &models.Episode{}
	utils.ParseBody(r, updateepisode)
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	ID, err := strconv.ParseInt(episodeId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")

	}
	episodeDetails, db := models.GetepisodeById(ID)
	if updateepisode.EpisodeNo != "" {
		episodeDetails.EpisodeNo = updateepisode.EpisodeNo
	}
	if updateepisode.EpisodeName != "" {
		episodeDetails.EpisodeName = updateepisode.EpisodeName
	}
	if updateepisode.Description != "" {
		episodeDetails.Description = updateepisode.Description
	}
	if updateepisode.SName != "" {
		episodeDetails.SName = updateepisode.SName
	}
	if updateepisode.Platform != "" {
		episodeDetails.Platform = updateepisode.Platform
	}
	db.Save(&episodeDetails)
	res, _ := json.Marshal(episodeDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Season

var Newseason models.Season

func Getseason(w http.ResponseWriter, r *http.Request) {
	limit, er := strconv.Atoi(r.URL.Query()["limit"][0])
	page, er1 := strconv.Atoi(r.URL.Query()["page"][0])

	if er != nil || er1 != nil {
		fmt.Println("Error")
	}
	var paginate = paginations.Pagination{}
	paginate.Limit = limit
	paginate.Page = page
	season, err := models.GetAllseason(paginate)
	if season == nil || err != nil {
		w.Write([]byte("No movies found or Error Happened"))
	}
	res, _ := json.Marshal(season)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetseasonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seasonId := vars["seasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	seasonDetails, _ := models.GetseasonById(ID)
	res, _ := json.Marshal(seasonDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Createseason(w http.ResponseWriter, r *http.Request) {
	Createseason := &models.Season{}
	utils.ParseBody(r, Createseason)
	b := Createseason.Createseason()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Deleteseason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seasonId := vars["eseasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	season := models.Deleteseason(ID)
	res, _ := json.Marshal(season)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Updateseason(w http.ResponseWriter, r *http.Request) {
	var updateseason = &models.Season{}
	utils.ParseBody(r, updateseason)
	vars := mux.Vars(r)
	seasonId := vars["seasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")

	}
	seasonDetails, db := models.GetseasonById(ID)
	if updateseason.SeasonNo != "" {
		seasonDetails.SeasonNo = updateseason.SeasonNo
	}
	if updateseason.SeasonName != "" {
		seasonDetails.SeasonName = updateseason.SeasonName
	}
	if updateseason.ShowName != "" {
		seasonDetails.ShowName = updateseason.ShowName
	}
	if updateseason.Episodes != "" {
		seasonDetails.Episodes = updateseason.Episodes
	}

	db.Save(&seasonDetails)
	res, _ := json.Marshal(seasonDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
