package repositories

import (
	"database/sql"
	"log"
	"netflixReboot/main/masters/models"
)

//FilmRepositoryImpl app
type FilmRepositoryImpl struct {
	db *sql.DB
}

//SelectAllFilm app
func (s FilmRepositoryImpl) SelectAllFilm() ([]*models.Films, error) {
	data, err := s.db.Query("SELECT * FROM film")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var result = []*models.Films{}
	for data.Next() {
		var film = models.Films{}
		var err = data.Scan(&film.ID, &film.Title, &film.Duration, &film.ImageURL, &film.Synopsis)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, &film)
	}
	if err = data.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

//SelectFilm app
func (s FilmRepositoryImpl) SelectFilm(id int) (*models.Films, error) {
	var film = new(models.Films)
	err := s.db.QueryRow("SELECT * FROM film WHERE id = ?", id).Scan(&film.ID, &film.Title, &film.Duration, &film.ImageURL, &film.Synopsis)
	if err != nil {
		return nil, err
	}
	return film, nil
}

//AddFilm app
func (s FilmRepositoryImpl) AddFilm(inFilm *models.Films) error {

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "INSERT INTO film(title, duration, image_url, synopsis) values (?,?,?,?);"
	_, err = tx.Exec(query, inFilm.Title, inFilm.Duration, inFilm.ImageURL, inFilm.Synopsis)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

//InitFilmRepositoryImpl app
func InitFilmRepositoryImpl(db *sql.DB) FilmRepository {
	return &FilmRepositoryImpl{db}
}
