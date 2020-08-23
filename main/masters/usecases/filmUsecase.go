package usecases

import (
	"netflixReboot/main/masters/models"
)

//FilmUsecase app
type FilmUsecase interface {
	GetAllFilm() ([]*models.Films, error)
	GetFilm(id int) (*models.Films, error)
	PostFilm(inFilm *models.Films) error
}
