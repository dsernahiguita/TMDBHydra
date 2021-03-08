/**
* Errors
* Handling the errors from the complete application
* @author  Diana Lucia Serna Higuita
 */
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
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
	ErrorEnterTVSerie                   /* 8: Error when the user is entering the tvSerie  */
	ErrorGetNextPage                    /* 9: Error when the user is entering the next page */
	ErrorSelectSerieId                  /* 10: Error when the user is entering the serie Id */
	ErrorSelectSeasonId                 /* 11: Error when the user is entering the season Id  */
	ErrorSelectEpisodeId                /* 12: Error when the use is entering the episode Id */
)

type LogRecord struct {
	Service    string
	Error_Type string
	Error_Id   ErrorId
	Error      string
	Timestamp  int64
}

const Service = "CliGo"

type ErrorMessage struct {
	ErrorId ErrorId
	Message string
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
	case ErrorEnterTVSerie:
		return "Error when the user is entering the tvSerie"
	case ErrorGetNextPage:
		return "Error when the user is entering the next page"
	case ErrorSelectSerieId:
		return "Error when the user is entering the serie Id"
	case ErrorSelectSeasonId:
		return "Error when the user is entering the season Id"
	case ErrorSelectEpisodeId:
		return "Error when the use is entering the episode Id"

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
