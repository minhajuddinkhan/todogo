package utils

import (
	"encoding/json"
	"net/http"

	boom "github.com/darahayes/go-boom"
)

//Respond Respond
func Respond(w http.ResponseWriter, i interface{}) {
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(i)
	if err != nil {
		boom.Internal(w, err)
		return
	}

}
