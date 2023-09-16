package main

import (
	. "api-mysql/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/movie", GetMovie)
	router.POST("/movie/create", PostMovie)
	router.PUT("/movie/:id/update", UpdateMovie)
	router.DELETE("/movie/:id/delete", DeleteMovie)

	router.GET("/age-rating-category", GetAgeRatingCategory)
	router.POST("/age-rating-category/create", PostAgeRatingCategory)
	router.PUT("/age-rating-category/:id/update", UpdateAgeRatingCategory)
	router.DELETE("/age-rating-category/:id/delete", DeleteAgeRatingCategory)

	fmt.Println("Server Running at Port 9898")
	log.Fatal(http.ListenAndServe(":9898", router))
}
