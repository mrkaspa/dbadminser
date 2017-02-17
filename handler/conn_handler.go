package handler

import (
	"net/http"

	"encoding/json"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/mrkaspa/dbadminser/logic"
	"github.com/mrkaspa/dbadminser/store"
)

type connHandler struct {
	connManager logic.ConnManager
	connStore   store.ConnStore
}

func (c connHandler) storeHandler(w http.ResponseWriter, r *http.Request) {
	var conn store.Conn
	err := json.NewDecoder(r.Body).Decode(&conn)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	err = c.connStore.StoreConn(conn)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c connHandler) deleteHandler(w http.ResponseWriter, r *http.Request) {
	connIDStr := mux.Vars(r)["id"]
	connID, _ := strconv.Atoi(connIDStr)
	err := c.connStore.DeleteConn(connID)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c connHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	conns := c.connStore.ListConn()
	sendOkJSON(w, conns)
}
