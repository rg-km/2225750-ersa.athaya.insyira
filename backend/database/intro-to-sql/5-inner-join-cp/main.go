package main

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db}
}

func (r *MovieRepository) FetchMovies() ([]model.Movie, error) {

	// Task : create a query to fetch all movies
	// 1. create a query to fetch all movies
	// 2. use inner join to fetch all movies and their genres and their directors

	// TODO: answer here
	sqlStmt := `
		SELECT 
			m.id,
			m.title,
			m.description,
			m.year,
			g.name AS genre_id,
			d.name AS director_id
		FROM movies m
		INNER JOIN genres g ON g.id = m.genre_id
		INNER JOIN directors d ON d.id = m.director_id	
	`

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []model.Movie
	for rows.Next() {
		var m model.Movie
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Description,
			&m.Year,
			&m.GenreName,
			&m.DirectorName,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

func (r *MovieRepository) FetchMovieByID(id int64) (*model.Movie, error) {

	// Task : create a query to fetch a movie by id
	// 1. create a query to fetch a movie by id
	// 2. use inner join to fetch all movies and their genres and their directors

	// TODO: answer here
	sqlStmt := `
		SELECT
			m.id,
			m.title,
			m.description,
			m.year,
			g.name AS genre_id,
			d.name AS director_name
		FROM movies m
		INNER JOIN genres g ON g.id = m.genre_id
		INNER JOIN directors d ON d.id = m.director_id
		WHERE m.id = ?
	`

	row := r.db.QueryRow(sqlStmt, id)

	var m model.Movie
	err := row.Scan(
		&m.ID,
		&m.Title,
		&m.Description,
		&m.Year,
		&m.GenreName,
		&m.DirectorName,
	)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
