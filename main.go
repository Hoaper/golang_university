package main

import (
	"context"
	"github.com/Hoaper/golang_university/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func main() {
	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	router := mux.NewRouter()
	routes.SetRoutes(router, client)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // You can specify allowed origins here
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	err = http.ListenAndServe(":5000", corsHandler(router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
