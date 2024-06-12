package models

type Movie struct {
	Id          *int64  `json:"id"`
	DirectorID  *int64  `json:"director_id"`
	Title       *string `json:"title"`
	ReleaseYear *int32  `json:"release_year"`
	MusicBy     *int64  `json:"music_by"`
	WrittenBy   *int64  `json:"written_by"`
}
