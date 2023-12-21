package main

import (
    "fmt"
    "net/http"
    "time"
    "workout-web.kristerley/internal/data"
)

func (app *application) createExerciseHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w,"Creates an exercise")
}

func (app *application) showExerciseHandler(w http.ResponseWriter, r *http.Request){
    id, err := app.readIdParams(r)
    if err != nil || id<1 {
        http.NotFound(w,r)
        return
    }

    exercise := data.Exercise{
        ID: id,
        CreatedAt: time.Now(),
        Name: "Air Squat",
        MainBodyPart: "Quadriceps",
        SecondaryBodyParts: []string{"Glutes", "Abdominals"},
        Purpose: "Strength",
        Version: 1,
    }

    err = app.writeJSON(w,http.StatusOK, envelope{"exercise":exercise}, nil)
    if err != nil {
        app.logger.Println(err)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
    }

}

    
