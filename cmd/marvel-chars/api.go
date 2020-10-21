package main

import (
	"log"
	"net/http"
)

// GetCharacters returns all marvel characters
func GetCharacters(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCharacters")
	w.WriteHeader(http.StatusOK)
}

// GetCharacterByID returns marvel character
// given its ID
func GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCharacterByID")
	w.WriteHeader(http.StatusOK)
}
