package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"netflixReboot/main/masters/models"
	"netflixReboot/main/masters/usecases"
	"netflixReboot/utils"
	"strconv"

	"github.com/gorilla/mux"
)

//FilmHandler app
type FilmHandler struct {
	FilmUsecase usecases.FilmUsecase
}

//FilmController app
func FilmController(r *mux.Router, s usecases.FilmUsecase) {
	filmHandler := FilmHandler{s}
	film := r.PathPrefix("/film").Subrouter()

	//Get
	film.HandleFunc("", filmHandler.ListFilm).Methods(http.MethodGet)
	film.HandleFunc("/{id}", filmHandler.Film).Methods(http.MethodGet)
	//Post
	film.HandleFunc("", filmHandler.PostFilm).Methods(http.MethodPost)

}

//ListFilm app
func (s *FilmHandler) ListFilm(w http.ResponseWriter, r *http.Request) {
	films, err := s.FilmUsecase.GetAllFilm()
	var filmResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		filmResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", Data: err.Error()}
		utils.ResponseWrite(&filmResponse, w)
		log.Println(err)
	} else {
		filmResponse = utils.Response{Status: http.StatusOK, Message: "Get All Film Success", Data: films}
		utils.ResponseWrite(&filmResponse, w)
	}
	log.Println("Endpoint hit: Get All Films")
}

//Film app
func (s *FilmHandler) Film(w http.ResponseWriter, r *http.Request) {
	var filmResponse utils.Response
	ex := mux.Vars(r)
	idINT, err := strconv.Atoi(ex["id"])
	film, err := s.FilmUsecase.GetFilm(idINT)
	w.Header().Set("content-type", "application/json")
	if err != nil {
		filmResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", Data: err.Error()}
		utils.ResponseWrite(&filmResponse, w)
		log.Println(err)
	} else {
		filmResponse = utils.Response{Status: http.StatusOK, Message: "Get film Success", Data: film}
		utils.ResponseWrite(&filmResponse, w)
	}
	log.Println("Endpoint hit: Get film")
}

//PostFilm app
func (s *FilmHandler) PostFilm(w http.ResponseWriter, r *http.Request) {
	var inFilm models.Films
	var filmResponse utils.Response
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&inFilm)
	if err != nil {
		log.Println(err)
		// w.Write([]byte("cant read JSON"))
	}
	err = s.FilmUsecase.PostFilm(&inFilm)
	if err != nil {
		filmResponse = utils.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		utils.ResponseWrite(&filmResponse, w)
		log.Println(err)
	} else {
		filmResponse = utils.Response{Status: http.StatusAccepted, Message: "Post film Success", Data: inFilm}
		utils.ResponseWrite(&filmResponse, w)
	}
	log.Println("Endpoint hit: Post film")
}
