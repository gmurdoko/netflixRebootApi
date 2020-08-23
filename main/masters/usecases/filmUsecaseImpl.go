package usecases

import (
	"netflixReboot/main/masters/models"
	"netflixReboot/main/masters/repositories"
	"netflixReboot/utils/validation"
)

//FilmUsecaseImpl spp
type FilmUsecaseImpl struct {
	filmRepo repositories.FilmRepository
}

//GetAllFilm app
func (s FilmUsecaseImpl) GetAllFilm() ([]*models.Films, error) {
	films, err := s.filmRepo.SelectAllFilm()
	if err != nil {
		return nil, err
	}
	return films, nil
}

//GetFilm app
func (s FilmUsecaseImpl) GetFilm(id int) (*models.Films, error) {
	film, err := s.filmRepo.SelectFilm(id)
	if err != nil {
		return nil, err
	}
	return film, nil
}

//PostFilm app
func (s FilmUsecaseImpl) PostFilm(inFilm *models.Films) error {
	err := validation.ValidateInputNotNil(inFilm.Title, inFilm.Duration, inFilm.ImageURL, inFilm.Synopsis)
	if err != nil {
		return err
	}
	err = s.filmRepo.AddFilm(inFilm)
	if err != nil {
		return err
	}
	return nil
}

//InitFilmUsecaseImpl app
func InitFilmUsecaseImpl(filmRepo repositories.FilmRepository) FilmUsecase {
	return &FilmUsecaseImpl{filmRepo}
}
