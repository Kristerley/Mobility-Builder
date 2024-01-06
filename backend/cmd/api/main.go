package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
    // "workout-web.kristerley/internal/dbCon"
    //
	"context"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

const version = "1.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
	client *mongo.Client
}

//type client struct  {
// 	client *mongo.Client

// }

// var client *mongo.Client

func main() {
	var cfg config
	var client *mongo.Client

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	client = dbConnect()
	defer client.Disconnect(context.TODO())

	app := &application{
		config: cfg,
		logger: logger,
		client: client,
	}


	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}



	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
