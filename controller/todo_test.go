package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-todo/schemas"
	"net/http"
	"net/http/httptest"
	"testing"
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

			fmt.Println((*req).Body)

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
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/todo?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			TC.GetTodoListHandler(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}

func TestGetTodoByIDHandler(t *testing.T) {}

func TestCompleteTodoHandler(t *testing.T) {}

func TestGetDeleteTodoHandler(t *testing.T) {}
