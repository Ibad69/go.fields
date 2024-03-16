package api

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ibad69/go.fields/pkg/players/transport"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DbHost string
	DbPort int
	DbUser string
	DbPw   string
	DbName string

	AppHost string
	AppPort string
}

func Start(cfg *Config) {
	// start the server and everything else initialize the router and activate the routes
	// initialize chi router
	// initialize the database connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.DbHost))
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	r := chi.NewRouter()
	transport.Activate(r, client)
	err = http.ListenAndServe(":"+cfg.AppPort, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("Server running on port: ", cfg.AppPort)
}
