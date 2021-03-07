module github.com/TMDBHydra/BackedForFrontend

replace github.com/TMDBHydra/BackedForFrontend/pkg/errors => ./pkg/errors

replace github.com/TMDBHydra/BackedForFrontend/pkg/config => ./pkg/config

replace github.com/TMDBHydra/BackedForFrontend/pkg/api => ./pkg/api

require (
	github.com/TMDBHydra/BackedForFrontend/pkg/api v0.0.0
	github.com/TMDBHydra/BackedForFrontend/pkg/config v0.0.0
	github.com/TMDBHydra/BackedForFrontend/pkg/errors v0.0.0
	github.com/gorilla/mux v1.7.4
	github.com/rs/cors v1.7.0 // indirect
)

go 1.14
