package main

import (
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Home ...дефолтный обработчик
func (app *application) Home(w http.ResponseWriter, _ *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies app and running",
		Version: "1.0.0",
	}
	out, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	// добавляем заголовки
	w.Header().Set("Content-Type", "applications/json")
	// добавляем статус код
	w.WriteHeader(http.StatusOK)
	// возвращаем ответ
	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
	}
}

// AllMovies ...обработчик возвращает список всех фильмов
func (app *application) AllMovies(w http.ResponseWriter, _ *http.Request) {
	// слайс фильмов
	movies := make([]models.Movie, 2, 10)
	// 1 фильм
	releaseDate, _ := time.Parse("2006-01-02", "1986-03-07")
	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: releaseDate,
		MPAARating:  "R",
		RunTime:     116,
		Description: "A very nice movie",
		CreateAt:    time.Now(),
		UpdateFiend: time.Now(),
	}
	movies[highlander.ID-1] = highlander
	// 1 фильм
	releaseDate, _ = time.Parse("2006-01-02", "1981-06-12")
	rotla := models.Movie{
		ID:          2,
		Title:       "Raiders of the Lost Ark",
		ReleaseDate: releaseDate,
		MPAARating:  "PG-13",
		RunTime:     115,
		Description: "Another very nice movie",
		CreateAt:    time.Now(),
		UpdateFiend: time.Now(),
	}
	movies[rotla.ID-1] = rotla

	// создаем json
	out, err := json.Marshal(movies)
	if err != nil {
		log.Println(err)
	}
	// добавляем заголовки
	w.Header().Set("Content-Type", "applications/json")
	// добавляем статус код
	w.WriteHeader(http.StatusOK)
	// возвращаем ответ
	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
	}
}
