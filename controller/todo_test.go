package controller_test

import (
	"bytes"
	"encoding/json"
	"go-todo/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostTodoHandler(t *testing.T) {

	// テスト用のリクエストを作成
	reqModel := models.Todo{
		Task: "test-todo1",
	}
	reqBody, _ := json.Marshal(reqModel)                                           // 構造体をbyte型のjsonに変換
	req := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(reqBody)) // byte.NewBufferは引数に与えられたバイトを含むバッファを作成

	// recorderはレスポンスの検証機能を提供する。(responseWriterの代わりに使用)
	recorder := httptest.NewRecorder() // NewRecorderはhttp.ResponseWriterインターフェースを満たす。
	TC.PostTodoHandler(recorder, req)

	// レスポンスのステータスコードを検証
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 得られたレスポンスを構造体に変換
	var gotRes models.Todo
	json.Unmarshal(recorder.Body.Bytes(), &gotRes) // Unmarshalはバイトを構造体に流し込む

	// 得られたレスポンスとレクエストが一致するか検証
	const expectedTask = "test-todo1"
	if gotRes.Task != expectedTask {
		t.Errorf("handler returned unexpected task: got %s want %s", gotRes.Task, expectedTask)
	}

}
