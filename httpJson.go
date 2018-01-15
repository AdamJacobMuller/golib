package golib

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func MarshalAndWriteJson(w http.ResponseWriter, object interface{}) {
	jsonDocument, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(503)
		log.Errorln(err)
		return
	}
	w.WriteHeader(200)
	w.Header().Add("Content-type", "application/json")
	w.Write(jsonDocument)
}
