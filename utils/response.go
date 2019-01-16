package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON ...
type ResponseJSON struct {
	Error   interface{} `json:"error,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// Response this function will give Content-Type JSON on given response
func Response(res http.ResponseWriter, error interface{}, message interface{}, data interface{}, httpCode int, meta interface{}) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(httpCode)

	response := &ResponseJSON{
		Error:   error,
		Data:    data,
		Message: message,
		Meta:    meta,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		panic(err)
	}
}
