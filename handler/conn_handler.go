package handler

import (
	"net/http"

	"github.com/mrkaspa/dbadminser/logic"
	"github.com/mrkaspa/dbadminser/store"
	"encoding/json"
)

type connHandler struct {
	connManager logic.ConnManager
	connStore   store.ConnStore
}

func (c connHandler) storeHandler(w http.ResponseWriter, r *http.Request) {
	conn := store.Conn{}
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
