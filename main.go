package main

import (
	"log"
	"net/http"

	"example.com/go-api-test/server"

	"go.uber.org/fx"
)

func initApp() *fx.App {
	app := fx.New(
		fx.Provide(server.SetupRouter),

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

	if err := app.Start(nil); err != nil {
		log.Fatal(err)
	}
}
