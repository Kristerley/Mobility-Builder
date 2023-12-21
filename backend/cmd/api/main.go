package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "workout-web.kristerley/internal/keys"
    
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

const version = "1.0"

type config struct{
    port int
    env string
}

type application struct {
    config config
    logger *log.Logger
}

func main(){
    var cfg config 

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(keys.MONGODB_URI).SetServerAPIOptions(serverAPI)

    client, errDB := mongo.Connect(context.TODO(), opts)
    if errDB != nil {
        panic(errDB)
    }


    defer func() {
        if errDB = client.Disconnect(context.TODO()); errDB != nil {
          panic(errDB)
        }
    }() 

    if errDB := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); errDB != nil {
        panic(errDB)
     }
     fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

    flag.IntVar(&cfg.port, "port", 4000, "API server port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
    flag.Parse()

    logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

    app := &application{
        config: cfg,
        logger: logger,
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

    srv := &http.Server {
        Addr: fmt.Sprintf(":%d", cfg.port),
        Handler: app.routes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
    err := srv.ListenAndServe()
    logger.Fatal(err)


}

