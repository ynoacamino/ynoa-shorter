package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ynoacamino/ynoa-shorter/db"
)

func CreateShorter(w http.ResponseWriter, r *http.Request) {
	var url db.Url

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	query, err := db.Query.CreateShorter(r.Context(), db.CreateShorterParams{
		ShortUrl:    url.ShortUrl,
		OriginalUrl: url.OriginalUrl,
		UserID:      url.UserID,
		Public:      url.Public,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(query)
}

func GetPublicShorters(w http.ResponseWriter, r *http.Request) {
	query, err := db.Query.GetPublicShorters(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(query)
}

func GetPrivateShorters(w http.ResponseWriter, r *http.Request) {
	var url db.Url

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if url.UserID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User ID is required"))
		return
	}

	query, err := db.Query.GetPrivateShorters(r.Context(), url.UserID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(query)
}

func DeleteShorter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	param, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	urlID := int32(param)

	query, err := db.Query.DeleteShorter(r.Context(), urlID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(query)
}

func UpdateShorter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	param, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	urlID := int32(param)

	var url db.Url

	err = json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	query, err := db.Query.UpdateShorter(r.Context(), db.UpdateShorterParams{
		UrlID:    urlID,
		ShortUrl: url.ShortUrl,
		Public:   url.Public,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(query)
}

func SetUpShorterRoutes(router *mux.Router) {
	router.HandleFunc("/public", GetPublicShorters).Methods("GET")
	router.HandleFunc("/", GetPrivateShorters).Methods("GET")
	router.HandleFunc("/", CreateShorter).Methods("POST")
	router.HandleFunc("/{id}", UpdateShorter).Methods("PUT")
	router.HandleFunc("/{id}", DeleteShorter).Methods("DELETE")
}
