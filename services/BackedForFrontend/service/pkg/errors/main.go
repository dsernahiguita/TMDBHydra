/**
* Errors
* Handling the errors from the complete application
* @author  Diana Lucia Serna Higuita
 */
package errors

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ErrorType uint

const (
	NoType ErrorType = iota
	Fatal
	Error
	Warning
	HttpRequest
)

type ErrorId uint

const (
	NoErrorId                   ErrorId = iota
	ErrorRequestBodyBadlyFormed         /* 1: Body of the file request (post) badly formed */
	ErrorRequestParameterEmpty          /* 2: One of the parameters of the request is empty */
	ErrorConfigFileNotFound             /* 3: Config file not found */
	ErrorConfigFileUnreadable           /* 4: Config file not readable */
	ErrorGetTVSerie                     /* 5: Error by trying to get the tvseries */
	ErrorGetSeasons                     /* 6: Error by trying to get the tvserie's seasons */
	ErrorGetEpisodes                    /* 7: Error by trying to get the season's episodes */
)

type LogRecord struct {
	Service    string
	Error_Type string
	Error_Id   ErrorId
	Error      string
	Timestamp  int64
}

const Service = "BackendForFrontend"

type ErrorMessage struct {
	ErrorId ErrorId
	Message string
}

type Config struct {
	DatabaseLogs           string `json:"databaseLogs"`
	DatabaseLogsCollection string `json:"databaseLogsCollection"`
	DatabaseHost           string `json:"databaseHost"`
	DatabaseUser           string `json:"databaseUser"`
	DatabasePassword       string `json:"databasePassword"`
}

var DatabaseLogs string
var DatabaseLogsCollection string
var DatabaseHost string
var DatabaseUser string
var DatabasePassword string

/**
* Load database credentials
* @return bool
 */
func LoadDatabaseCredentials() bool {
	/* Read config file to get the credential of the logs database */
	jsonFile, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		return false
	}
	var config Config

	/* we unmarshal our byteArray which contains our
	jsonFile's content into 'config' which we defined above */
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		return false
	}
	DatabaseLogs = config.DatabaseLogs
	DatabaseLogsCollection = config.DatabaseLogsCollection
	DatabaseHost = config.DatabaseHost
	DatabaseUser = config.DatabaseUser
	DatabasePassword = config.DatabasePassword

	return true
}

/**
* Handling errors Http request
* @param http.ResponseWriter w
* @param string errorMessage
* @param errorId ErrorId
* @param bool logErrors
 */
func HandlingErrorsHttpRequest(w http.ResponseWriter, errorMessage string, errorId ErrorId, logErrors bool) {
	if logErrors {
		context := fmt.Sprintf("%v %v", errorId.getErrorIdDescription(), errorMessage)
		HttpRequest.Log(context, errorId)
	}
	errorIdDescription := fmt.Sprintf("%s. %s", errorId.getErrorIdDescription(), errorMessage)
	messageJson, err := json.Marshal(&ErrorMessage{ErrorId: errorId, Message: errorIdDescription})
	if err != nil {
	}
	if errorId == ErrorRequestBodyBadlyFormed || errorId == ErrorRequestParameterEmpty {
		http.Error(w, string(messageJson), http.StatusBadRequest)
		return
	} else {
		http.Error(w, string(messageJson), http.StatusInternalServerError)
		return
	}
}

/**
* Handling errors
* Print the error
* if the variable logErrors is true, then the errors will be log in the file log/####
* if the ErrorType = Fatal the programm will be abort
* @param ErrorType errorType
* @param error error
* @param bool logErrors
* @param ErrorId errorId
 */
func (errorType ErrorType) HandlingErrors(err error, logErrors bool, errorId ErrorId) {
	errorTypeName := errorType.getErrorTypeName()
	errorIdDescription := errorId.getErrorIdDescription()
	description := errorTypeName + " " + errorIdDescription
	if err != nil {
		description = description + ": " + err.Error()
	}
	currentTime := time.Now()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d - %s\n", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), description)
	if logErrors {
		errorType.Log(description, errorId)
	}
	switch errorType {
	case Fatal:
		os.Exit(1)
	}
}

/**
* Log
* log errors in a file
* @param ErrorType errorType
* @param string context
* @return bool
 */
func (errorType ErrorType) Log(error string, errorId ErrorId) bool {
	/* read date time creation file */
	currentTime := time.Now()
	pwd, _ := os.Getwd()
	/* create folder log if not exist */
	folder := pwd + "/logs"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, 0755)
	}
	nameFile := folder + "/log_" + fmt.Sprintf("%d%02d%02d", currentTime.Year(), currentTime.Month(), currentTime.Day())

	f, err := os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	errorTypeName := errorType.getErrorTypeName()
	content := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d, %v, %v, %v\n", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), errorTypeName, errorId, error)
	if _, err := f.Write([]byte(content)); err != nil {
		fmt.Println(err)
		return false
	}
	f.Close()
	errorType.saveLogDB(error, errorId)

	return true
}

/**
* Save log DB
* @param errorType ErrorType
* @param string error
* @param ErrorId errorId
* @param bool
 */
func (errorType ErrorType) saveLogDB(error string, errorId ErrorId) bool {
	/* set client options */
	mongoDBServerConnection := fmt.Sprintf(
		"mongodb+srv://%v:%v@%v/?retryWrites=true&w=majority",
		DatabaseUser,
		DatabasePassword,
		DatabaseHost,
	)

	clientOptions := options.Client().ApplyURI(mongoDBServerConnection)
	/* Connect to MongoDB */
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return false
	}

	/* Check the connection */
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	logCollection := client.Database(DatabaseLogs).Collection(DatabaseLogsCollection)
	errorTypeName := errorType.getErrorTypeName()
	var logRecord LogRecord
	logRecord.Service = Service
	logRecord.Error_Type = errorTypeName
	logRecord.Error_Id = errorId
	logRecord.Error = error
	logRecord.Timestamp = time.Now().Unix()
	_, err = logCollection.InsertOne(context.TODO(), logRecord)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/**
* Get error id description
* get description of the error id
* @param ErrorId errorId
* @return string errorIdDescription
 */
func (errorId ErrorId) getErrorIdDescription() string {
	switch errorId {
	case NoErrorId:
		return "NoId"
	case ErrorRequestBodyBadlyFormed:
		return "Please check the body of your post request, it is badly formed"
	case ErrorRequestParameterEmpty:
		return "Please check the body of your request, one or more parameters are empty"
	case ErrorConfigFileNotFound:
		return "The config file is not founded"
	case ErrorConfigFileUnreadable:
		return "The config file is not readable"
	case ErrorGetTVSerie:
		return "Error by trying to get the tvseries"
	case ErrorGetSeasons:
		return "Error by trying to get the tvserie's seasons"
	case ErrorGetEpisodes:
		return "Error by trying to get the season's episodes"
	default:
		return "NoId"
	}
}

/**
* Get error type name
* get name of type error
* @param errorType ErrorType
* @return string errorTypeName
 */
func (errorType ErrorType) getErrorTypeName() string {
	switch errorType {
	case NoType:
		return "NoType"
	case Fatal:
		return "Fatal"
	case Error:
		return "Error"
	case Warning:
		return "Warning"
	case HttpRequest:
		return "HttpRequest"
	default:
		return "NoType"
	}
}
