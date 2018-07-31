package utils

import (
	"encoding/json"
	"net/http"
)

//DecodeRequest DecodeRequest
func DecodeRequest(r *http.Request, model interface{}) error {

	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&model)

}
