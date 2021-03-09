# Cli GO (Command Line Tool)
Command line tool to simulate user ui interaction and data flow
Steps:
1. When the cli tool is started, the user is asked for a tv series title (free text)
2. The user enters a title and presses enter a list of matched series is shown.
3. The user can move between pages by typing in the number of the page he/she wants to view.
4. If the user doesn't want to move between pages, he/she can enter press the key enter
5. The user is asked to pick a serie title from that result list.
6. The list of seasons of the series are shown.
7. The use is asked to pick a season from the result list.
8. The list of episodes of the season are shown.
9. The use is asked to pick a episode from the result list.
10. In the end the user gets displayed a title and a summary of the chosen episode..
This service consumes the service BackedForFrontend

## Installation
### Requirements
* Go
* Config file "config.json": this file must be located in the same folder where
the application was installed, path: config/config.json.
This file allow to set the following options
  * logErrors boolean: the logs will be stored in txt files
  located into the path logs/
  * backendTMDBHydra: backend service


### Run Manual
* Go to the path TMDBHydra/services/BackedForFrontend/service
* Run
```
Go run main.go
```

## Unit tests
### Run unit test for package without dependencies
* Go to the path TMDBHydra/services/CliGo/service/pkg/errors
* Run
```
go test -v
```

### Run unit test for package with dependencies
The packages Api and Config are tested here
* Go to the path TMDBHydra/services/CliGo/service
* Run
```
go test -v
```

## Error handling:
The application errors are classified:
* Fatal: the application will be closed
* Error: the application will not stop and the errors will be stored in Mongo DB
* Warning: the application will not stop and the warnings will be stored in Mongo DB

### Response errors
* Code 1: Error, Please check the body of your post request, it is badly formed
* Code 2: Error, Please check the body of your post request, one or more of the parameters
are empty
* Code 3: Error, Error the config file is not founded
* Code 4: Error, Error the config file is not readable
* Code 5: Error, Error by trying to get the tvseries
* Code 6: Error, Error by trying to get the tvserie's seasons
* Code 7: Error, Error by trying to get the season's episodes
* Code 8:  Error when the user is entering the tvSerie
* Code 9: Error when the user is entering the next page
* Code 10: Error when the user is entering the serie Id
* Code 11: Error when the user is entering the season Id
* Code 12: Error when the use is entering the episode Id


### Logs
## License
