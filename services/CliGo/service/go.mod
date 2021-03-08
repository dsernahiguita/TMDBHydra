module github.com/TMDBHydra/CliGo

replace github.com/TMDBHydra/CliGo/pkg/errors => ./pkg/errors

replace github.com/TMDBHydra/CliGo/pkg/config => ./pkg/config

replace github.com/TMDBHydra/CliGo/pkg/api => ./pkg/api

require (
	github.com/TMDBHydra/CliGo/pkg/api v0.0.0
	github.com/TMDBHydra/CliGo/pkg/config v0.0.0
	github.com/TMDBHydra/CliGo/pkg/errors v0.0.0
)

go 1.14
