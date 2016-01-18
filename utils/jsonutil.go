package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJson(w http.ResponseWriter, JsonType interface{}) {
	js, err := json.Marshal(JsonType)
	if err != nil {
		RespondError(w, err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetJson(body []byte) (jsonSource interface{}) {
	err := json.Unmarshal(body, &jsonSource)
	if err != nil {
		log.Print("template executing error: ", err)
	}
	return
}
