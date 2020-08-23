package repositories

import (
	"netflixReboot/main/masters/models"
)

//FilmRepository app
type FilmRepository interface {
	SelectAllFilm() ([]*models.Films, error)
	SelectFilm(id int) (*models.Films, error)
	AddFilm(inFilm *models.Films) error
}
