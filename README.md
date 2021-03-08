# TMDBHydra
## Story/Assignment
As a mercedes benz S class customer, I want that my children
are able to watchtv series (in the build-in backseat screens) during long journeys, that i can focus on driving.

## POC:
This POC shall simulate the user ui interaction and data flow via a command line tool.
When the cli tool is started, the user is asked for a tv series title (free text), After the user enters a title and presses enter a list of matched series is shown.
The user has to pick a series title from that result list. In the next step the user has choose from a season and an episode.
In the end the user gets displayed a title and a summary of the chosen episode.##

## Restrictions
* use golang
* use tmdb api as provider for tv informations https://developers.themoviedb.org/3/getting-started/introduction
* You have to register and aquire an api-key (service is free of charge)
* Please provide us the source code as a public accessible git repo (github, gitlab,...)

## Architecture
The system is composed of three components:

### BackedForFrontend
Rest API that implements the following services:
* Get TV series
* Get seasons
* Get Episodes
This API consumes the The Movie Database API

located: services/BackedForFrontend

After the Installation this service is available under the Port: 4060

### WebInterfaceReact
Web application that allows the user in a graphical interface to searches series using free text, select seasons and episodes.

located: services/WebInterfaceReact

After the Installation this service is available under the Port: 3001

### CliGo
Command line tool to simulate user ui interaction and data flow

located: services/CliGo

```
-------------------------------------    -------------------------------------
|                                    |   |                                   |
|       WEB Interface React          |   |    Cli Go (Commander line tool)   |  
|                                    |   |                                   |
-------------------------------------    -------------------------------------

----------------------------------------------------------------------------
|                                                                          |
|                          BackedForFrontend Api(Port 4060)                |
|                                                                          |   
----------------------------------------------------------------------------

----------------------------------------------------------------------------
|                                                                          |
|                                   TMDB Api                               |
|                                                                          |  
----------------------------------------------------------------------------
```

## Installation
### Requirements
Docker

1. Clone project
```
git clone git@github.com:dsernahiguita/TMDBHydra.git
```
2. Go to the path TMDBHydra/scripts/
3. Run
```
docker-compose build
```
4. Run
```
docker-compose up
```

### Logs
## License
