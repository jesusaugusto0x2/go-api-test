package server

import (
	"fmt"
	"net/http"

	"example.com/go-api-test/ent"
	"example.com/go-api-test/repository"
	"example.com/go-api-test/service"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func SetupRouter(lc fx.Lifecycle, client *ent.Client) http.Handler {
	r := chi.NewRouter()

	userRepo := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	r.Get(RouteUsers, userHandler.GetUsers)
	r.Post(RouteUsers, userHandler.CreateUser)
	r.Get(RouteUsersWithIDParam, userHandler.GetUser)
	r.Patch(RouteUsersWithIDParam, userHandler.UpdateUser)
	r.Delete(RouteUsersWithIDParam, userHandler.DeleteUser)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	return r
}
