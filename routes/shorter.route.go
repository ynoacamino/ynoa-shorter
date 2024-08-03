package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ynoacamino/ynoa-shorter/db"
	"github.com/ynoacamino/ynoa-shorter/models"
)

func CreateShorter(w http.ResponseWriter, r *http.Request) {
	var link models.Link
	json.NewDecoder(r.Body).Decode(&link)

	if !isLink(link) {
		throwBadRequest(w, errors.New("invalid link"))
		return
	}

	createdUser := db.DB.Create(&link)

	err := createdUser.Error
	if err != nil {
		throwBadRequest(w, err)
		return
	}

	json.NewEncoder(w).Encode(&link)
}

func GetPublicShorters(w http.ResponseWriter, r *http.Request) {
	var links []models.Link

	query := db.DB.Where(&models.Link{Public: true}).Find(&links)
	err := query.Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&links)
}

func GetPrivateShorters(w http.ResponseWriter, r *http.Request) {
	var links []models.Link

	var userLink struct {
		UserId string `json:"userId"`
	}

	json.NewDecoder(r.Body).Decode(&userLink)

	query := db.DB.Where(&models.Link{UserId: userLink.UserId}).Find(&links)
	err := query.Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&links)
}

func DeleteShorter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	linkId := params["id"]

	if linkId == "" {
		throwBadRequest(w, errors.New("invalid link"))
		return
	}

	var link models.Link
	query := db.DB.Unscoped().Delete(&link, linkId)
	err := query.Error
	if err != nil {
		throwBadRequest(w, err)
		return
	}

	json.NewEncoder(w).Encode(&link)
}

func UpdateShorter(w http.ResponseWriter, r *http.Request) {
	var link models.Link
	json.NewDecoder(r.Body).Decode(&link)

	if !isLink(link) {
		throwBadRequest(w, errors.New("invalid link"))
		return
	}

	query := db.DB.Save(&link)
	err := query.Error
	if err != nil {
		throwBadRequest(w, err)
		return
	}

	json.NewEncoder(w).Encode(&link)
}

func SetUpShorterRoutes(router *mux.Router) {
	router.HandleFunc("/link/public", GetPublicShorters).Methods("GET")
	router.HandleFunc("/link", GetPrivateShorters).Methods("GET")
	router.HandleFunc("/link", CreateShorter).Methods("POST")
	router.HandleFunc("/link/{id}", UpdateShorter).Methods("PUT")
	router.HandleFunc("/link/{id}", DeleteShorter).Methods("DELETE")
}

func isLink(link models.Link) bool {
	if len(link.Real) < 4 {
		return false
	}

	if len(link.UserId) == 0 {
		return false
	}

	return true
}

func throwBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
