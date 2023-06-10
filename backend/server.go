package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  // インスタンスを作成
  e := echo.New()

  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
	e.Use(middleware.CORS())

  // ルートを設定
  e.GET("/", hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する
	e.GET("/foods", getFoods)

  // サーバーをポート番号1323で起動
  e.Logger.Fatal(e.Start(":1323"))
}

type Food struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func getFoods(c echo.Context) error {
	foods := []Food{
		{Name: "チョコレート", Price: 100},
		{Name: "ポテトチップス", Price: 200},
		{Name: "コーラ", Price: 150},
		{Name: "アイス", Price: 200},
	}
	return c.JSON(http.StatusOK, foods)
}
