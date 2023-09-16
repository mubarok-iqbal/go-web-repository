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

func GetMovie(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	movieRepository := repository.MovieRepository(config.MySQL())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	movies, err := movieRepository.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, movies, http.StatusOK)
}

func PostMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	movieRepository := repository.MovieRepository(config.MySQL())

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.Movie
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := movieRepository.Insert(ctx, mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	movieRepository := repository.MovieRepository(config.MySQL())

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.Movie

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMovie = ps.ByName("id")

	if err := movieRepository.Update(ctx, mov, idMovie); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteMovie(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	movieRepository := repository.MovieRepository(config.MySQL())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMovie = ps.ByName("id")
	if err := movieRepository.Delete(ctx, idMovie); err != nil {
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
