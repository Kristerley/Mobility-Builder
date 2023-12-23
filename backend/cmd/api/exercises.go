package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "context"
    // "time"
    "workout-web.kristerley/internal/data"
    "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"

)

// func (app *application) createExerciseHandler(w http.ResponseWriter, r *http.Request){
//     fmt.Fprintln(w,"Creates an exercise")
// }

func (app *application) createExerciseHandler(w http.ResponseWriter, r *http.Request){
    var input struct{
        NameEn string
        NameRu string
        MainBodyPart string
        SecondaryBodyParts []string
        Purpose string
    }

    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w,"Decode Error", http.StatusInternalServerError)
        return
    }
    fmt.Fprint(w,"%+v\n",input)
}

func (app *application) showExerciseHandler(w http.ResponseWriter, r *http.Request){
    exercises := []data.Exercise{}
    // var input struct{
    //     Purpose string
    // }
    // err := json.NewDecoder(r.Body).Decode(&input)

    coll := app.client.Database("Exercises").Collection("Mobility")

    cursor, err :=  coll.Find(context.TODO(), bson.D{})
    // results := []bson{}
    if err != nil { app.logger.Fatal(err) }
    defer cursor.Close(context.Background())

    err = cursor.All(context.TODO(), &exercises)
    if err != nil {
        app.logger.Println(err)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
    }
    err = app.writeJSON(w,http.StatusOK, exercises, nil)
}

    
