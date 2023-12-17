package models

import "time"

type Movie struct {
	// уникальный идентификатор фильм
	ID int `json:"id"`
	// название фильма
	Title string `json:"title"`
	// дата выпуска (релиза) фильма
	ReleaseDate time.Time `json:"release_date"`
	// длительность фильма
	RunTime int `json:"runTime"`
	// рейтинг фильма, установленный Motion Picture Association of America (MPAA), например PG-13
	MPAARating string `json:"mpaa_rating"`
	// описание фильма
	Description string `json:"description"`
	// ссылка на изображение, представляющее фильм
	Image string `json:"image"`
	// время создания записи о фильме
	CreateAt time.Time `json:"create_at"`
	// время последнего обновления записи о фильме
	UpdateFiend time.Time `json:"update_fiend"`
}
