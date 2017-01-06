package handler

import (
    "encoding/json"
    "net/http"

    "github.com/mrkaspa/geoserver/utils"
)

func sendOkJSON(w http.ResponseWriter, data interface{}) {
    sendJSONWithStatus(w, data, http.StatusOK)
}

func sendJSONWithStatus(w http.ResponseWriter, data interface{}, code int) {
    utils.Log.Infof("Response %d = %v", code, data)
    w.WriteHeader(code)
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(data)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
    }
}
