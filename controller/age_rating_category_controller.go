package controller

import (
	"api-mysql/config"
	"api-mysql/models"
	"api-mysql/repository"
	"api-mysql/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Read
// GetAgeRatingCategory
func GetAgeRatingCategory(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	ageRatingCategoryRepository := repository.AgeRatingCategoryRepository(config.MySQL())

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	ratings, err := ageRatingCategoryRepository.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, ratings, http.StatusOK)
}

// Create
// PostAgeRatingCategory
func PostAgeRatingCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ageRatingCategoryRepository := repository.AgeRatingCategoryRepository(config.MySQL())

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var rating models.AgeRatingCategory

	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := ageRatingCategoryRepository.Insert(ctx, rating); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateAgeRatingCategory
func UpdateAgeRatingCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ageRatingCategoryRepository := repository.AgeRatingCategoryRepository(config.MySQL())

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var rating models.AgeRatingCategory

	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idRating = ps.ByName("id")

	if err := ageRatingCategoryRepository.Update(ctx, rating, idRating); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteAgeRatingCategory
func DeleteAgeRatingCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ageRatingCategoryRepository := repository.AgeRatingCategoryRepository(config.MySQL())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idRating = ps.ByName("id")

	if err := ageRatingCategoryRepository.Delete(ctx, idRating); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}
