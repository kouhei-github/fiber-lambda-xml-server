package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kouhei-github/fiber-sample-framework/route"
	"os"
)

func main() {
	app := fiber.New()

	// ルーターの設定
	router := &route.Router{FiberApp: app}

	// CORS (Cross Origin Resource Sharing)の設定
	// アクセスを許可するドメイン等を設定します
	router.FiberApp.Use(cors.New(cors.Config{AllowHeaders: "Origin, Content-Type, Accept"}))

	route.LoadRouter(router)

	// Webサーバー起動時のエラーハンドリング => localhostの時コメントイン必要
	if os.Getenv("ENVIRONMENT") == "local" {
		if err := app.Listen(":8080"); err != nil {
			panic(err)
		}
	}

	// AWS Lambdaとの連携設定
	//lambda.Start(httpadapter.NewV2(handler).ProxyWithContext)
}
