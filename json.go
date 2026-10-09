package main

import (
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

}
