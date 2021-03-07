# Backend for frontend service
Rest API server to implement the following services:
* Get TV series
* Get seasons
* Get Episode


## Installation
### Requirements
* Config file "config.json": this file must be located in the same folder where
the application was installed, path: config/config.json.
This file allow to set the following options
  * logErrors boolean: the logs will be stored in the mongo DB and in txt files
  located into the path logs/
  * portRestAPI: port exposed used by the endpoints
  * databaseLogs: Mongo DB where the errors logs are stored
  * databaseLogsCollection: collection from the database where the errors are stored
  * databaseHost
  * databaseUser
  * databasePassword
  * backendServiceTMDB: backend service The move DB provider for tv informations
  * apiKeyTMDB

### Installation Manual

### Installation using dockers

* Run the Dockerfile
```
docker build ./
```
* Start the container
```

```

## Usage
### API Endpoints
* Get tv series
Note: the parameter page is optional, by default the page 1 is returned
```
GET localhost:4060/api/frontend/tvserie
Body:
{
  "query": "Modern Family",
  "page": 1
}
```
Response:
```
{
  "page": 1,
  "total_pages": 1,
  "total_results": 2,
  "results": [
    {
      "id": 1421,
      "name": "Modern Family",
      "original_name": "Modern Family",
      "overview": "The Pritchett-Dunphy-Tucker clan is a wonderfully large and blended family. They give us an honest and often hilarious look into the sometimes warm, sometimes twisted, embrace of the modern family."
    },
    {
        "id": 30509,
        "name": "The Madness of Modern Families",
        "original_name": "The Madness of Modern Families",
        "overview": "Light-hearted look at the absurd behaviour displayed by British parents desperate to get it right for their offspring."
    }
  ]
}
```

* Get seasons
```
GET localhost:4060/api/frontend/seasons
Body:
{
  "tvserieId": 1421
}
```
Response:
```
{
  "name": "Modern Family",
  "number_of_seasons": 11,
  "number_of_episodes": 250,
  "seasons": [
    {
      "id": 147409,
      "name": "Specials",
      "overview": "",
      "season_number": 0
    },
    {
       "id": 3751,
       "name": "Season 1",
       "overview": "Modern Family takes a refreshing and funny view of what it means to raise a family in this hectic day and age.  Multi-cultural relationships, adoption, and same-sex marriage are just a few of the timely issues faced by the show’s three wildly-diverse broods.  No matter the size or shape, family always comes first in this hilariously “modern” look at life, love, and laughter.",
       "season_number": 1
   },
   ....
  ]
}
```

* Get episodes:
```
POST localhost:4060/api/frontend/episodes
Body:
{
  "tvserieId": 1421
  "season": 1
}
```
Response:
```
{

}
```


## Error handling:
The application errors are classified:
* Fatal: the application will be closed
* Error: the application will not stop and the errors will be stored in Mongo DB
* HttpRequest: the application will not stop the errors will be sent into the
  http response and stored in the Mongo DB
  The response has the following structure:
  ```
  {
    ErrorId: <optional, only if the request had errors>,
    Message: <required: if there is an error, then the description of the error is
              placed in this field. If the request was executed successfully then in the message
              is the description of the executed process>
  }
  ```
* Warning: the application will not stop and the warnings will be stored in Mongo DB


### Response errors
* Code 1: Error, Please check the body of your post request, it is badly formed
* Code 2: Error, Please check the body of your post request, one or more of the parameters
are empty
* Code 3: Error, Error the config file is not founded
* Code 4: Error, Error the config file is not readable


### Logs
## License
