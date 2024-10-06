package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-todo/schemas"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestPostTodoHandler(t *testing.T) {

	var tests = []struct {
		name       string
		req        schemas.PostTodoRequest
		resultCode int
	}{
		{name: "正常系", req: schemas.PostTodoRequest{Task: "aaa"}, resultCode: http.StatusOK},
		{name: "空のタスク", req: schemas.PostTodoRequest{Task: ""}, resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.req)                                             // 構造体をbyte型のjsonに変換
			req := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(body)) // byte.NewBufferは引数に与えられたバイトを含むバッファを作成

			// recorderはレスポンスの検証機能を提供する。(responseWriterの代わりに使用)
			rec := httptest.NewRecorder() // NewRecorderはhttp.ResponseWriterインターフェースを満たす。
			TC.PostTodoHandler(rec, req)

			if rec.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, rec.Code)
			}
		})
	}

}

func TestGetTodoListHandler(t *testing.T) {

	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "正常系", query: "1", resultCode: http.StatusOK},
		{name: "異常系", query: "aaa", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/todo?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			rec := httptest.NewRecorder()

			TC.GetTodoListHandler(rec, req)

			if rec.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, rec.Code)
			}
		})
	}
}

func TestGetTodoByIDHandler(t *testing.T) {
	var tests = []struct {
		name       string
		id         string
		resultCode int
	}{
		{name: "正常系", id: "1", resultCode: http.StatusOK},
		{name: "異常系", id: "5", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/todo/"+tt.id, nil)

			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todo/{id:[0-9]+}", TC.GetTodoByIDHandler).Methods(http.MethodGet)
			r.ServeHTTP(rec, req)

			if rec.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, rec.Code)
			}
		})
	}
}

func TestCompleteTodoHandler(t *testing.T) {
	var tests = []struct {
		name       string
		id         string
		resultCode int
	}{
		{name: "正常系", id: "1", resultCode: http.StatusOK},
		{name: "異常系", id: "5", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPatch, "/todo/"+tt.id, nil)

			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todo/{id:[0-9]+}", TC.CompleteTodoHandler).Methods(http.MethodPatch)
			r.ServeHTTP(rec, req)

			if rec.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, rec.Code)
			}
		})
	}
}

func TestGetDeleteTodoHandler(t *testing.T) {
	var tests = []struct {
		name       string
		id         string
		resultCode int
	}{
		{name: "正常系", id: "1", resultCode: http.StatusOK},
		{name: "異常系", id: "5", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/todo/"+tt.id, nil)

			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todo/{id:[0-9]+}", TC.DeleteTodoHandler).Methods(http.MethodDelete)
			r.ServeHTTP(rec, req)

			if rec.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, rec.Code)
			}
		})
	}
}
