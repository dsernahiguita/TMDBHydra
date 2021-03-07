/**
* Api
* @author  Diana Lucia Serna Higuita
 */

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	Config "github.com/TMDBHydra/BackedForFrontend/pkg/config"
	Errors "github.com/TMDBHydra/BackedForFrontend/pkg/errors"
)

type bodyGetTVSerie struct {
	Query string
	Page  int
}

type bodyGetSeasons struct {
	TVSerieId int
}

type bodyGetEpisodes struct {
	TVSerieId int
	Season    int
}

type bodyGetSummaryEpisode struct {
	Episode string
}

/**
* Get
* Processing Get calls
* @param http.ResponseWriter w
* @param http.Request r
 */
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called: this call will be not proccessed"}`))
}

/**
* Post
* Processing Post calls
* @param http.ResponseWriter w
* @param http.Request r
 */
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called: this call will be not proccessed"}`))
}

/**
* Put
* Processing Put calls
* @param http.ResponseWriter w
* @param http.Request r
 */
func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called: this call will be not proccessed"}`))
}

/**
* Delete
* Processing Delete calls
* @param http.ResponseWriter w
* @param http.Request r
 */
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called: this call will be not proccessed"}`))
}

/**
* Get tv series
*
* @param http.ResponseWriter w
* @param http.Request r
 */
func getTVSeries(w http.ResponseWriter, r *http.Request) {
	var p bodyGetTVSerie
	/* Try to decode the request body into the struct. If there is an error,
	respond to the client with the error message and a 400 status code.*/
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorRequestBodyBadlyFormed, Config.LogErrors)
		return
	}
	/* if one of the parameters in empty we should return error message 400 */
	if len(p.Query) == 0 {
		Errors.HandlingErrorsHttpRequest(w, "", Errors.ErrorRequestParameterEmpty, Config.LogErrors)
		return
	}
	query := p.Query
	var page int
	if p.Page == 0 {
		page = 1
	} else {
		page = p.Page
	}

	results, err := GetTVSeries(query, page)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorGetTVSerie, Config.LogErrors)
		return
	}

	/* send response {page, total_pages, total_results, results: [{id, name, origial_name, overview}]}*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message, err := json.Marshal(results)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorGetTVSerie, Config.LogErrors)
		return
	}
	w.Write([]byte(message))
}

/**
* Get seasons
*
* @param http.ResponseWriter w
* @param http.Request r
 */
func getSeasons(w http.ResponseWriter, r *http.Request) {
	var p bodyGetSeasons
	/* Try to decode the request body into the struct. If there is an error,
	respond to the client with the error message and a 400 status code.*/
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorRequestBodyBadlyFormed, Config.LogErrors)
		return
	}
	/* if one of the parameters in empty we should return error message 400 */
	if p.TVSerieId == 0 {
		Errors.HandlingErrorsHttpRequest(w, "", Errors.ErrorRequestParameterEmpty, Config.LogErrors)
		return
	}
	tvSerieId := p.TVSerieId

	results, err := GetSeasons(tvSerieId)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorGetSeasons, Config.LogErrors)
		return
	}

	/* send response {name, number_of_seasons, number_of_episodes, seasons: [{id, name, overview, season_number}]}*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message, err := json.Marshal(results)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorGetSeasons, Config.LogErrors)
		return
	}
	w.Write([]byte(message))
}

/**
* Get episodes
*
* @param http.ResponseWriter w
* @param http.Request r
 */
func getEpisodes(w http.ResponseWriter, r *http.Request) {
	var p bodyGetEpisodes
	/* Try to decode the request body into the struct. If there is an error,
	respond to the client with the error message and a 400 status code.*/
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorRequestBodyBadlyFormed, Config.LogErrors)
		return
	}
	/* if one of the parameters in empty we should return error message 400 */
	if p.TVSerieId == 0 || p.Season == 0 {
		Errors.HandlingErrorsHttpRequest(w, "", Errors.ErrorRequestParameterEmpty, Config.LogErrors)
		return
	}
	tvSerieId := p.TVSerieId
	season := p.Season

	results, err := GetEpisodes(tvSerieId, season)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorGetEpisodes, Config.LogErrors)
		return
	}

	/* send response {id, name, overview, season_number, episodes: [{id, name, overview, episode_number}]} */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message, err := json.Marshal(results)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorGetEpisodes, Config.LogErrors)
		return
	}
	w.Write([]byte(message))
}

/**
* Get summary episode
*
* @param http.ResponseWriter w
* @param http.Request r
 */
func getSummaryEpisode(w http.ResponseWriter, r *http.Request) {
	var p bodyGetSummaryEpisode
	/* Try to decode the request body into the struct. If there is an error,
	respond to the client with the error message and a 400 status code.*/
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, err.Error(), Errors.ErrorRequestBodyBadlyFormed, Config.LogErrors)
		return
	}
	/* if one of the parameters in empty we should return error message 400 */
	if len(p.Episode) == 0 {
		Errors.HandlingErrorsHttpRequest(w, "", Errors.ErrorRequestParameterEmpty, Config.LogErrors)
		return
	}
	episode := p.Episode

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	message := fmt.Sprintf("Episode %s ", episode)
	w.Write([]byte(`{"message": "` + message + `", "episode": "` + episode + `"}`))
}

/**
* API
* Here are implemented a request router and dispatcher for matching incoming request
* that includes the format /storeDataset
* Handler implemented:
* POST /storeDataset/rx store a RX file into IPFS
* POST /storeDataset/ct store serveral CT files into IPFS
* The port used by the rest API is defined into the config.json file
 */
func Main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/frontend").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)

	/* Call service encrypt file */
	api.HandleFunc("/tvserie", getTVSeries).Methods(http.MethodGet)
	api.HandleFunc("/seasons", getSeasons).Methods(http.MethodGet)
	api.HandleFunc("/episodes", getEpisodes).Methods(http.MethodGet)
	api.HandleFunc("/episode", getSummaryEpisode).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+Config.PortRestAPI, r))
}
