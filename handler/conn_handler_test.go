package handler

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/golang/mock/gomock"
    "github.com/mrkaspa/dbadminser/store"
    "github.com/mrkaspa/dbadminser/store/mock"
    "github.com/stretchr/testify/assert"
)

func getMockStore(t *testing.T) (store.ConnStore, func()) {
    ctrl := gomock.NewController(t)
    mockConnStore := mock.NewMockConnStore(ctrl)
    mockConnStore.EXPECT().StoreConn(gomock.Any()).Return(nil)
    return mockConnStore, ctrl.Finish
}

func TestConnHandler_storeHandler(t *testing.T) {
    t.Run("when the conn is saved", whenTheConnIsSaved)
}

func whenTheConnIsSaved(t *testing.T)  {
    mockConnStore, finish := getMockStore(t)
    defer finish()
    connHandlerTest := connHandler{connStore: mockConnStore}
    conn := store.Conn{}
    data, _ := json.Marshal(conn)
    req, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(data))
    w := httptest.NewRecorder()
    connHandlerTest.storeHandler(w, req)
    assert.Equal(t, w.Code, http.StatusOK)
}
