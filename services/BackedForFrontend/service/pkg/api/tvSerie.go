package api

import (
	"encoding/json"
	"errors"
	"fmt"
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
* @return array seasons
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
		fmt.Println(tvSerieDetails)
		if err != nil {
			return tvSerieDetails, err
		}

		return tvSerieDetails, nil
	}

	bodyString := string(body)
	err = errors.New(bodyString)
	return tvSerieDetails, err
}
