package main

import (
    "fmt"
    "workout-web.kristerley/internal/keys"
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// type mClient struct {
//     client *mongo.Client
// }
var client *mongo.Client

func dbConnect() (*mongo.Client) {
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(keys.MONGODB_URI).SetServerAPIOptions(serverAPI)

    client, errDB := mongo.Connect(context.TODO(), opts)
    if errDB != nil {
        panic(errDB)
    }



    if errDB := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); errDB != nil {
        panic(errDB)
     }
     fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
    // coll := client.Database("Exercises").Collection("Mobility")
    // result, err := coll.InsertOne(context.TODO(), bson.D{{"Dude","Guy"}})
    // fmt.Println(err)
    // fmt.Println(result)

     return client
}
