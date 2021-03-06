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

/**
* Get tv serie
* @param string query
* @param int page
* @return BodyGetTVSeries tvSeries
* @return error err
 */
func GetTVSerie(query string, page int) (BodyGetTVSeries, error) {
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
