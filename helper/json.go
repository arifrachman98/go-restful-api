package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, res interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(res)
	PanicHelper(err)
}

func WriteToResponseBody(w http.ResponseWriter, r interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(r)
	PanicHelper(err)
}
