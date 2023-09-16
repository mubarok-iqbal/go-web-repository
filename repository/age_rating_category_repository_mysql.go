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

type ageRatingCategoryRepositoryMysql struct {
	DB *sql.DB
}

const (
	ageRatingCategoryTable          = "age_rating_category"
	ageRatingCategoryLayoutDateTime = "2006-01-02 15:04:05"
)

// GetAll AgeRatingCategory
func (repo *ageRatingCategoryRepositoryMysql) GetAll(ctx context.Context) ([]models.AgeRatingCategory, error) {

	var ratings []models.AgeRatingCategory

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", ageRatingCategoryTable)

	rowQuery, err := repo.DB.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var rating models.AgeRatingCategory
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&rating.ID,
			&rating.Name,
			&rating.Description,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		rating.CreatedAt, err = time.Parse(ageRatingCategoryLayoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		rating.UpdatedAt, err = time.Parse(ageRatingCategoryLayoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		ratings = append(ratings, rating)
	}

	return ratings, nil
}

// Insert AgeRatingCategory
func (repo *ageRatingCategoryRepositoryMysql) Insert(ctx context.Context, rating models.AgeRatingCategory) error {

	queryText := fmt.Sprintf("INSERT INTO %v (name, description, created_at, updated_at) values('%v','%v', NOW(), NOW())", ageRatingCategoryTable,
		rating.Name,
		rating.Description)

	_, err := repo.DB.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update AgeRatingCategory
func (repo *ageRatingCategoryRepositoryMysql) Update(ctx context.Context, rating models.AgeRatingCategory, id string) error {

	queryText := fmt.Sprintf("UPDATE %v set name ='%s', description = '%s', updated_at = NOW() where id = %s",
		ageRatingCategoryTable,
		rating.Name,
		rating.Description,
		id,
	)
	fmt.Println(queryText)

	_, err := repo.DB.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete AgeRatingCategory
func (repo *ageRatingCategoryRepositoryMysql) Delete(ctx context.Context, id string) error {

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", ageRatingCategoryTable, id)

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
