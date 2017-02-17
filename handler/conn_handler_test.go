package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrkaspa/dbadminser/store"
	"github.com/stretchr/testify/assert"
)

func getMockStore(t *testing.T) (*store.MockConnStore, func()) {
	ctrl := gomock.NewController(t)
	mockConnStore := store.NewMockConnStore(ctrl)
	return mockConnStore, ctrl.Finish
}

func getMockStoreSave(t *testing.T, ret interface{}) (store.ConnStore, func()) {
	ctrl := gomock.NewController(t)
	mockConnStore := store.NewMockConnStore(ctrl)
	mockConnStore.EXPECT().StoreConn(gomock.Any()).Return(ret)
	return mockConnStore, ctrl.Finish
}

func getMockStoreDelete(t *testing.T, ret interface{}) (store.ConnStore, func()) {
	ctrl := gomock.NewController(t)
	mockConnStore := store.NewMockConnStore(ctrl)
	mockConnStore.EXPECT().DeleteConn(gomock.Any()).Return(ret)
	return mockConnStore, ctrl.Finish
}

func TestConnHandler_storeHandler(t *testing.T) {
	t.Run("when the conn is saved", whenTheConnIsSaved)
	t.Run("when the conn is not saved", whenTheConnIsNotSaved)
}

func whenTheConnIsSaved(t *testing.T) {
	mockConnStore, finish := getMockStore(t)
	mockConnStore.EXPECT().StoreConn(gomock.Any()).Return(nil)
	defer finish()
	connHandlerTest := connHandler{connStore: mockConnStore}
	conn := store.Conn{}
	data, _ := json.Marshal(conn)
	fmt.Println(string(data))
	req, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(data))
	w := httptest.NewRecorder()
	connHandlerTest.storeHandler(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}

func whenTheConnIsNotSaved(t *testing.T) {
	mockConnStore, finish := getMockStore(t)
	mockConnStore.EXPECT().StoreConn(gomock.Any()).Return(errors.New("Conn not saved"))
	defer finish()
	connHandlerTest := connHandler{connStore: mockConnStore}
	conn := store.Conn{}
	data, _ := json.Marshal(conn)
	req, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(data))
	w := httptest.NewRecorder()
	connHandlerTest.storeHandler(w, req)
	assert.Equal(t, w.Code, http.StatusNotAcceptable)
}

func TestConnHandler_deleteHandler(t *testing.T) {
	mockConnStore, finish := getMockStore(t)
	mockConnStore.EXPECT().DeleteConn(gomock.Any()).Return(nil)
	defer finish()
	connHandlerTest := connHandler{connStore: mockConnStore}
	req, _ := http.NewRequest(http.MethodDelete, "", nil)
	w := httptest.NewRecorder()
	connHandlerTest.deleteHandler(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestConnHandler_indexHandler(t *testing.T) {
	mockConnStore, finish := getMockStore(t)
	conn := store.Conn{}
	mockConnStore.EXPECT().ListConn().Return([]store.Conn{conn})
	defer finish()
	connHandlerTest := connHandler{connStore: mockConnStore}
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	w := httptest.NewRecorder()
	connHandlerTest.indexHandler(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	var conns []store.Conn
	json.NewDecoder(bytes.NewReader(w.Body.Bytes())).Decode(&conns)
	assert.Len(t, conns, 1)
}
