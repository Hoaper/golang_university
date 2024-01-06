package routes

import (
	"github.com/Hoaper/golang_university/app/handlers"
	"github.com/Hoaper/golang_university/app/repositories"
	"github.com/Hoaper/golang_university/app/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func SetRoutes(router *mux.Router, client *mongo.Client) {
	authHandler := handlers.NewAuthHandler(
		services.NewAuthService(
			repositories.NewUserRepository(client),
		),
	)

	router.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", authHandler.LogoutHandler).Methods("GET")
	router.HandleFunc("/auth/register", authHandler.RegisterHandler).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
}
