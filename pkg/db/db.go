package db

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/DuvanAlbarracin/movies_movies/pkg/models"
	"github.com/DuvanAlbarracin/movies_movies/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Conn *pgxpool.Pool
}

func Init(url string) Handler {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalln("Error creating connection with the database:", err)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalln("Error while acquiring connection from the database pool:", err)
	}
	defer conn.Release()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalln("Error pinging the database:", err)
	}

	log.Println("Database conection success!")

	createMoviesTable(pool)
	createGenrerTable(pool)
	createGenrerMovieTable(pool)
	populateGenreTable()

	return Handler{Conn: pool}
}

func createMoviesTable(pool *pgxpool.Pool) (err error) {
	_, err = pool.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS movies (id SERIAL PRIMARY KEY, title VARCHAR(50) UNIQUE NOT NULL, release_year INTEGER NOT NULL, director_id INTEGER REFERENCES profiles (id), music_by VARCHAR(30), written_by VARCHAR(30));")
	if err != nil {
		log.Fatalln("Error creating the Movies table")
		return
	}

	log.Println("Movies table created succesfully!")
	return nil
}

func createGenrerTable(pool *pgxpool.Pool) (err error) {
	_, err = pool.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS genrers (id SERIAL PRIMARY KEY, name VARCHAR(30) UNIQUE NOT NULL);")
	if err != nil {
		log.Fatalln("Error creating the Genres table")
		return
	}

	log.Println("Genres table created succesfully!")
	return nil
}

func populateGenreTable() (err error) {
	file, err := os.Open("../../genres.csv")
	if err != nil {
		log.Fatalln("Error opening the Genres csv seed")
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	genres, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Error reading the Genres csv seed")
		return
	}

	fmt.Println("Readed from csv: ", genres)
	return nil
}

func createGenrerMovieTable(pool *pgxpool.Pool) (err error) {
	_, err = pool.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS genremovie (genrer_id INT REFERENCES genders(id), movie_id INT REFERENCES movies(id), (genrer_id, movie_id) PRIMARY KEY );")
	if err != nil {
		log.Fatalln("Error creating the Joing table (GenreMovie)")
		return
	}

	log.Println("Join table (GenrerMovies) created succesfully!")
	return nil
}

func CreateMovie(pool *pgxpool.Pool, movie *models.Movie) (err error) {
	_, err = pool.Exec(context.Background(),
		"insert into movies(title, release_year, director_id, music_by, written_by) values($1, $2, $3, $4, $5)",
		movie.Title,
		movie.ReleaseYear,
		movie.DirectorID,
		movie.MusicBy,
		movie.WrittenBy,
	)
	return
}

func DeleteMovie(pool *pgxpool.Pool, movie_id int64) (err error) {
	_, err = pool.Exec(context.Background(),
		"DELETE FROM movies WHERE id = $1",
		movie_id,
	)
	return
}

func ModifyMovie(pool *pgxpool.Pool, id int64, movieMap map[string]any, args []any) (err error) {
	query, err := utils.BuildUpdateQueryString("movies", id, movieMap)
	_, err = pool.Exec(context.Background(), query, args...)
	return
}

func FindMovieById(pool *pgxpool.Pool, id int64) (models.Movie, error) {
	var movie models.Movie
	err := pool.QueryRow(context.Background(),
		"SELECT * FROM movies WHERE id = $1", id).Scan(&movie.Id, &movie.Title,
		&movie.ReleaseYear, &movie.DirectorID, &movie.MusicBy, &movie.WrittenBy,
	)
	return movie, err
}

func GetAllMovies(pool *pgxpool.Pool) ([]models.Movie, error) {
	var movies []models.Movie
	rows, _ := pool.Query(context.Background(),
		"SELECT * FROM movies ORDER BY title ASC")

	for rows.Next() {
		movie := models.Movie{}
		rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.ReleaseYear,
			&movie.DirectorID,
			&movie.MusicBy,
			&movie.WrittenBy,
		)
		movies = append(movies, movie)
	}
	err := rows.Err()

	return movies, err
}

func FindGenreById(pool *pgxpool.Pool, id int64) (models.Genre, error) {
	var genre models.Genre
	err := pool.QueryRow(context.Background(),
		"SELECT * FROM genres WHERE id = $1", id).Scan(&genre.Id, &genre.Name)
	return genre, err
}

func GetAllGenres(pool *pgxpool.Pool) ([]models.Genre, error) {
	var genres []models.Genre
	rows, _ := pool.Query(context.Background(),
		"SELECT * FROM genres ORDER BY name ASC")

	for rows.Next() {
		genre := models.Genre{}
		rows.Scan(
			&genre.Id,
			&genre.Name,
		)
		genres = append(genres, genre)
	}
	err := rows.Err()

	return genres, err
}

func AddGenderToMovie(pool *pgxpool.Pool, genre_id int64, movie_id int64) (err error) {
	_, err = pool.Exec(context.Background(),
		"insert into genremovie(genre_id, movie_id) values($1, $2)",
		genre_id,
		movie_id,
	)
	return
}

func RemoveGenderFromMovie(pool *pgxpool.Pool, genre_id int64, movie_id int64) (err error) {
	_, err = pool.Exec(context.Background(),
		"DELETE FROM genremovie WHERE genre_id = $1 AND movie_id = $2",
		genre_id,
		movie_id,
	)
	return
}
