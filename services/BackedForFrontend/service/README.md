# Backend for frontend service
Rest API server to implement the following services:
* Get TV series
* Get seasons episodes
* Get Episode


## Installation
### Requirements
* Config file "config.json": this file must be located in the same folder where
the application was installed, path: config/config.json.
This file allow to set the following options
  * logError boolean: the logs will be stored in the mongo DB
  * portRestAPI: port exposed used by the endpoints
  * databaseLogs: logs database
  * databaseLogsCollection: collection from the database where the errors will be stored

### Installation Manual

### Installation using dockers

* Run the Dockerfile
```
docker build ./
```
* Start the container
```
docker run -i -t -p 4042:4042 -v /my/local/path:/app/locaStorage
-v /my/local/path/keys:/app/vivamedKeys <imageID>
```


## Usage
### API Endpoints
* Get tv serie
```
GET localhost:4060/api/frontend/tvserie
Body:
{
  "title": "TEST"
}
```
Response:
```
{
    "message": ""
    "title": ""
}
```

* Get seasons episodes
```
GET localhost:4060/api/frontend/seasons
Body:
{
  "tvserie": ""
}
```
Response:
```
{
    "message": ""
    "tvserie": ""
}
```

* Get episode:
```
POST localhost:4060/api/frontend/episode
Body:
{
  "episode": ""
}
```
Response:
```
{
    "message": ""
    "episode": ""
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
