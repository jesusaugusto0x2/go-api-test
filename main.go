package main

import (
	"context"
	"log"
	"net/http"

	"example.com/go-api-test/config"
	"example.com/go-api-test/db"
	"example.com/go-api-test/server"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func initApp() *fx.App {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or could not be loaded; using system environment variables.")
	}

	app := fx.New(
		fx.Provide(
			config.NewConfig,
			db.NewEntClient,
			server.SetupRouter,
		),

		fx.Invoke(func(r http.Handler) {
			log.Println("Starting server on :8080")

			if err := http.ListenAndServe(":8080", r); err != nil {
				log.Fatal(err)
			}
		}),
	)

	return app
}

func main() {
	app := initApp()

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
