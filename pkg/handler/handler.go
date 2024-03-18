package handler

import (
	"net/http"
	"strconv"
	"encoding/json"
	"sync"


	"github.com/go-chi/chi"
	"github.com/TursunovImran/movie_rest_api_go"
	"github.com/jmoiron/sqlx"
)

var (
	actorsLock sync.RWMutex
	filmsLock  sync.RWMutex
)

func InitRouter(db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()
	
	r.Route("/actors", func(r chi.Router) {
		r.Get("/{actorID}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "actorID")
			actorID, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				http.Error(w, "Invalid actor ID", http.StatusBadRequest)
				return
			}
			getActor(w, r, db, actorID)
		})
	
		r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
			getAllActors(w, r, db)
		})
	
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			createActor(w, r, db)
		})
	
		r.Post("/change", func(w http.ResponseWriter, r *http.Request) {
			changeActor(w, r, db)
		})
	})
	
	r.Route("/films", func(r chi.Router) {
		r.Get("/{filmID}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "filmID")
			filmID, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				http.Error(w, "Invalid actor ID", http.StatusBadRequest)
				return
			}
			getFilm(w, r, db, filmID)
		})
	
		r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
			getAllFilms(w, r, db)
		})
	
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			createFilm(w, r, db)
		})
	
		r.Post("/change", func(w http.ResponseWriter, r *http.Request) {
			changeFilm(w, r, db)
		})
	})

	return r
}

func getActor(w http.ResponseWriter, r *http.Request, db *sqlx.DB, actorID int64) {
	actorsLock.RLock()
	defer actorsLock.RUnlock()

	row := db.QueryRow("SELECT name, gender, birthdate FROM actors WHERE id = $1", actorID)

	var name, gender, birthdate string

	err := row.Scan(&name, &gender, &birthdate)
	if err != nil {
		http.Error(w, "Failed to get actor information", http.StatusInternalServerError)
		return
	}

	actor := &movierestapigo.Actor{
		ID: int(actorID),
		Name: name,
		Gender: gender,
		Birthdate: birthdate,
	}
	
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(actor); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}
}

func createActor(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	actorsLock.Lock()
	defer actorsLock.Unlock()
}

func changeActor(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	actorsLock.Lock()
	defer actorsLock.Unlock()
}

func getAllActors(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	actorsLock.RLock()
	defer actorsLock.RUnlock()
}

func getFilm(w http.ResponseWriter, r *http.Request, db *sqlx.DB, filmID int64) {
	filmsLock.RLock()
	defer filmsLock.RUnlock()
}

func createFilm(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	filmsLock.Lock()
	defer filmsLock.Unlock()
}

func changeFilm(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	filmsLock.Lock()
	defer filmsLock.Unlock()
}

func getAllFilms(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	filmsLock.RLock()
	defer filmsLock.RUnlock()
}