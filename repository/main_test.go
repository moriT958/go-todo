package repository_test

import (
	"go-todo/repository/fixture"
	"log"
	"os"
	"testing"
)

var Fxt *fixture.Fixture

func TestMain(m *testing.M) {
	var err error

	// fixtureの用意
	Fxt, err = fixture.NewFixture()
	if err != nil {
		log.Println("Fail setting up test fixture.:", err)
		os.Exit(1)
	}

	// テストデータのセットアップ
	if err := Fxt.Setup(); err != nil {
		log.Println("Fail setting up test data.:", err)
		os.Exit(1)
	}

	// テスト実行
	m.Run()

	// 後片付け
	if err := Fxt.Teardown(); err != nil {
		log.Println("Fail teardonw.:", err)
	}
}
