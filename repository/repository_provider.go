package repository

import "database/sql"

func MovieRepository(db *sql.DB) MovieRepositoryInterface {
	return &movieRepositoryMysql{DB: db}
}

func AgeRatingCategoryRepository(db *sql.DB) AgeRatingCategoryRepositoryInterface {
	return &ageRatingCategoryRepositoryMysql{DB: db}
}
