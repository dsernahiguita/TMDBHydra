/**
* Api
* @author  Diana Lucia Serna Higuita
 */

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	Config "github.com/TMDBHydra/BackedForFrontend/pkg/config"
	Errors "github.com/TMDBHydra/BackedForFrontend/pkg/errors"
)

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
func GetTVSeries(w http.ResponseWriter, r *http.Request) {
	/* query and page*/
	parameters := r.URL.Query()

	/* if one of the parameters in empty we should return error message 400 */
	if len(parameters.Get("query")) == 0 {
		Errors.HandlingErrorsHttpRequest(w, "", Errors.ErrorRequestParameterEmpty, Config.LogErrors)
		return
	}
	query := parameters.Get("query")

	/* if the parameter page is empty or not number we should use by default 1 */
	page, err := strconv.Atoi(parameters.Get("page"))
	if err != nil {
		page = 1
	}

	results, err := TMDBGetTVSeries(query, page)
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
func GetSeasons(w http.ResponseWriter, r *http.Request) {
	/* tvSerieId */
	parameters := r.URL.Query()

	/* if the parameter tvSerieId is empty we should return an error */
	tvSerieId, err := strconv.Atoi(parameters.Get("tvSerieId"))
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, "", Errors.ErrorRequestParameterEmpty, Config.LogErrors)
		return
	}

	results, err := TMDBGetSeasons(tvSerieId)
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
func GetEpisodes(w http.ResponseWriter, r *http.Request) {
	/* tvSerieId and season */
	parameters := r.URL.Query()

	/* if the parameter tvSerieId is empty we should return an error */
	tvSerieId, err := strconv.Atoi(parameters.Get("tvSerieId"))
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, "", Errors.ErrorRequestParameterEmpty, Config.LogErrors)
		return
	}

	/* if the parameter season is empty we should return an error */
	season, err := strconv.Atoi(parameters.Get("season"))
	if err != nil {
		Errors.HandlingErrorsHttpRequest(w, "", Errors.ErrorRequestParameterEmpty, Config.LogErrors)
		return
	}

	results, err := TMDBGetEpisodes(tvSerieId, season)
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
* API
* Here are implemented a request router and dispatcher for matching incoming request
* that includes the format /api/frontend
* Handler implemented:
* Get tvserie
* Get seasons (of a tv serie)
* Get episodes (of a season)
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
	api.HandleFunc("/tvserie", GetTVSeries).Methods(http.MethodGet)
	api.HandleFunc("/seasons", GetSeasons).Methods(http.MethodGet)
	api.HandleFunc("/episodes", GetEpisodes).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+Config.PortRestAPI, r))
}
