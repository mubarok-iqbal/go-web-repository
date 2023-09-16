package repository

import (
	"api-mysql/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type movieRepositoryMysql struct {
	DB *sql.DB
}

const (
	movieTable          = "movie"
	movieLayoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Movie
func (repo *movieRepositoryMysql) GetAll(ctx context.Context) ([]models.Movie, error) {

	var movies []models.Movie

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", movieTable)

	rowQuery, err := repo.DB.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var movie models.Movie
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&movie.ID,
			&movie.Title,
			&movie.Year,
			&createdAt,
			&updatedAt,
			&movie.AgeRatingCategoryId); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		movie.CreatedAt, err = time.Parse(movieLayoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		movie.UpdatedAt, err = time.Parse(movieLayoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// Insert Movie
func (repo *movieRepositoryMysql) Insert(ctx context.Context, movie models.Movie) error {

	queryText := fmt.Sprintf("INSERT INTO %v (title, year, created_at, updated_at, age_rating_category_id) values('%v',%v, NOW(), NOW(), %v)", movieTable,
		movie.Title,
		movie.Year,
		movie.AgeRatingCategoryId)

	_, err := repo.DB.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Movie
func (repo *movieRepositoryMysql) Update(ctx context.Context, movie models.Movie, id string) error {

	queryText := fmt.Sprintf("UPDATE %v set title ='%s', year = %d, updated_at = NOW(), age_rating_category_id=%v where id = %s",
		movieTable,
		movie.Title,
		movie.Year,
		movie.AgeRatingCategoryId,
		id,
	)
	fmt.Println(queryText)

	_, err := repo.DB.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Movie
func (repo *movieRepositoryMysql) Delete(ctx context.Context, id string) error {

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", movieTable, id)

	s, err := repo.DB.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()

	fmt.Println(check)

	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
