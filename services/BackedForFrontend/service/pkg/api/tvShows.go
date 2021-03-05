package api

import (
	"fmt"

	Config "github.com/TMDBHydra/BackedForFrontend/pkg/config"
)

func GetTVShows() {
	fmt.Println(Config.BackendServiceTMDB)
}
