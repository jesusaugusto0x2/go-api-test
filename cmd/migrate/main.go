package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"example.com/go-api-test/config"
	"example.com/go-api-test/ent"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or could not be loaded; using system environment variables.")
	}

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Error creating config: %v", err)
	}

	client, err := ent.Open(dialect.Postgres, cfg.DSN)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	fmt.Println("Migrations applied successfully")
}
