package routes

import (
	"github.com/ManivelRamu/CMS-Module/cmd/pkg/controllers"
	"github.com/ManivelRamu/CMS-Module/cmd/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movies", middleware.VerifyLogin(controllers.CreateMovie)).Methods("POST")
	router.HandleFunc("/movies", middleware.VerifyLogin(controllers.GetMovie)).Methods("GET")
	//router.HandleFunc("/movies/pagination", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/movies/search/{attribute}/{value}", controllers.SearchMovie).Methods("GET")
	router.HandleFunc("/movies/{movieId}", middleware.VerifyLogin(controllers.GetMovieById)).Methods("GET")
	router.HandleFunc("/movies/{movieId}", middleware.VerifyLogin(controllers.UpdateMovie)).Methods("PUT")
	router.HandleFunc("/movies/{movieId}", middleware.VerifyLogin(controllers.DeleteMovie)).Methods("DELETE")

	//show routes
	router.HandleFunc("/show", middleware.VerifyLogin(controllers.Createshow)).Methods("POST")
	router.HandleFunc("/show", middleware.VerifyLogin(controllers.Getshow)).Methods("GET")
	router.HandleFunc("/show/{showId}", middleware.VerifyLogin(controllers.GetshowById)).Methods("GET")
	router.HandleFunc("/show/{showId}", middleware.VerifyLogin(controllers.UpdateShow)).Methods("PUT")
	router.HandleFunc("/show/{showId}", middleware.VerifyLogin(controllers.Deleteshow)).Methods("DELETE")

	//season routes
	router.HandleFunc("/season", middleware.VerifyLogin(controllers.Createseason)).Methods("POST")
	router.HandleFunc("/season", middleware.VerifyLogin(controllers.Getseason)).Methods("GET")
	router.HandleFunc("/season/{seasonId}", middleware.VerifyLogin(controllers.GetseasonById)).Methods("GET")
	router.HandleFunc("/season/{seasonId}", middleware.VerifyLogin(controllers.Updateseason)).Methods("PUT")
	router.HandleFunc("/season/{seasonId}", middleware.VerifyLogin(controllers.Deleteseason)).Methods("DELETE")

	//EpisodeRoutes
	router.HandleFunc("/episode", middleware.VerifyLogin(controllers.Createepisode)).Methods("POST")
	router.HandleFunc("/episode", middleware.VerifyLogin(controllers.Getepisode)).Methods("GET")
	router.HandleFunc("/episode/{episodeId}", middleware.VerifyLogin(controllers.GetepisodeById)).Methods("GET")
	router.HandleFunc("/episode/{episodeId}", middleware.VerifyLogin(controllers.Updateepisode)).Methods("PUT")
	router.HandleFunc("/episode/{episodeId}", middleware.VerifyLogin(controllers.Deleteepisode)).Methods("DELETE")

}
