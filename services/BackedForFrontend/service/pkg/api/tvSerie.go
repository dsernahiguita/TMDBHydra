/**
* TV Serie
* @author  Diana Lucia Serna Higuita
 */

package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	Config "github.com/TMDBHydra/BackedForFrontend/pkg/config"
)

type ResultsTVSeries struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	OriginalName string `json:"original_name"`
	Overview     string `json:"overview"`
}

type BodyGetTVSeries struct {
	Page         int               `json:"page"`
	TotalPages   int               `json:"total_pages"`
	TotalResults int               `json:"total_results"`
	Results      []ResultsTVSeries `json:"results"`
}

type Season struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	SeasonNumber int    `json:"season_number"`
}

type BodyGetTVSerieDetails struct {
	Name             string   `json:"name"`
	NumberOfSeasons  int      `json:"number_of_seasons"`
	NumberOfEpisodes int      `json:"number_of_episodes"`
	Seasons          []Season `json:"seasons"`
}

type Episode struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Overview      string `json:"overview"`
	EpisodeNumber int    `json:"episode_number"`
}

type EpisodesSeason struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Overview     string    `json:"overview"`
	SeasonNumber int       `json:"season_number"`
	Episodes     []Episode `json:"episodes"`
}

/**
* Get tv series
* @param string query
* @param int page
* @return BodyGetTVSeries tvSeries
* @return error err
 */
func GetTVSeries(query string, page int) (BodyGetTVSeries, error) {
	var tvSeries BodyGetTVSeries
	endPoint := "search/tv"

	request, err := http.NewRequest("GET", Config.BackendServiceTMDB+endPoint, nil)
	if err != nil {
		return tvSeries, err
	}
	/* add the query parameters */
	q := request.URL.Query()
	q.Add("api_key", Config.ApiKeyTMDB)
	q.Add("language", "en-US")
	q.Add("include_adult", "false")
	q.Add("query", query)
	q.Add("page", strconv.Itoa(page))
	request.URL.RawQuery = q.Encode()

	/* do call */
	response, err := http.Get(request.URL.String())
	if err != nil {
		return tvSeries, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tvSeries, err
	}

	if response.StatusCode == 200 {
		err = json.Unmarshal(body, &tvSeries)
		if err != nil {
			return tvSeries, err
		}

		return tvSeries, nil
	}

	bodyString := string(body)
	err = errors.New(bodyString)
	return tvSeries, err
}

/**
* Get Seasons
* @param int tvSerieId
* @return object tvSerieDetails: this object has the field seasons
* @return error err
 */
func GetSeasons(tvSerieId int) (BodyGetTVSerieDetails, error) {
	var tvSerieDetails BodyGetTVSerieDetails
	endPoint := "tv/" + strconv.Itoa(tvSerieId)

	request, err := http.NewRequest("GET", Config.BackendServiceTMDB+endPoint, nil)
	if err != nil {
		return tvSerieDetails, err
	}
	/* add the query parameters */
	q := request.URL.Query()
	q.Add("api_key", Config.ApiKeyTMDB)
	q.Add("language", "en-US")
	request.URL.RawQuery = q.Encode()

	/* do call */
	response, err := http.Get(request.URL.String())
	if err != nil {
		return tvSerieDetails, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tvSerieDetails, err
	}

	if response.StatusCode == 200 {
		err = json.Unmarshal(body, &tvSerieDetails)
		if err != nil {
			return tvSerieDetails, err
		}

		return tvSerieDetails, nil
	}

	bodyString := string(body)
	err = errors.New(bodyString)
	return tvSerieDetails, err
}

/**
* Get episodes
* @param int tvSerieId
* @param int season
* @return object EpisodesSeason: this object has the field episodes
* @return error err
 */
func GetEpisodes(tvSerieId int, season int) (EpisodesSeason, error) {
	var episodesSeason EpisodesSeason
	endPoint := "tv/" + strconv.Itoa(tvSerieId) + "/season/" + strconv.Itoa(season)

	request, err := http.NewRequest("GET", Config.BackendServiceTMDB+endPoint, nil)
	if err != nil {
		return episodesSeason, err
	}
	/* add the query parameters */
	q := request.URL.Query()
	q.Add("api_key", Config.ApiKeyTMDB)
	q.Add("language", "en-US")
	request.URL.RawQuery = q.Encode()

	/* do call */
	response, err := http.Get(request.URL.String())
	if err != nil {
		return episodesSeason, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return episodesSeason, err
	}

	if response.StatusCode == 200 {
		err = json.Unmarshal(body, &episodesSeason)
		if err != nil {
			return episodesSeason, err
		}

		return episodesSeason, nil
	}

	bodyString := string(body)
	err = errors.New(bodyString)
	return episodesSeason, err
}
