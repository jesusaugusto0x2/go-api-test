# go-api-test

Just created this small api using Golang to learn some basics :)

I opted to follow the commonly known `controller -> service -> repository` pattern, so it's pretty much intuitive how it goes.

This also only interacts with a single `User` table which has 2 columns only (`name` and `email`) so the idea is just to have a common CRUD that does implement the overall REST methods.

- The controllers are under the `/server` folder
- Services under `/service` folder
- And repositories under `/repository` folder

## How to run

- Run `go mod tidy` to install dependencies.
- Added a `.env.example` file, here you can just create a `.env` file with the same exact key-values to get the app running.
- I also put a `docker-compose.yml` file in where I declare single postgres database just in case, so first run `docker compose up` to get the db living under the `5433` port in your local. You can use any other custom database if you want, but be sure then to update the values on your `.env` file.
- I understand the migration process should be ran separately, so under the `cmd/migrate/main.go` file is a small script to migrate the only table we got to run the project, make sure to run it by doing `go run cmd/migrate/main.go`
- Finally, just do `go run main.go` and the project should be running!

## Packages used

- [Go-Chi](https://go-chi.io/) for routing
- [Fx](https://uber-go.github.io/fx/index.html) for dependency injection
- [Ent](https://entgo.io/) for the ORM
