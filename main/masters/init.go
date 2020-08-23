package masters

import (
	"database/sql"
	"netflixReboot/main/masters/controllers"
	"netflixReboot/main/masters/repositories"
	"netflixReboot/main/masters/usecases"

	"github.com/gorilla/mux"
)

//Init app
func Init(r *mux.Router, db *sql.DB) {
	//Film
	filmRepo := repositories.InitFilmRepositoryImpl(db)
	filmUsecase := usecases.InitFilmUsecaseImpl(filmRepo)
	controllers.FilmController(r, filmUsecase)
}
