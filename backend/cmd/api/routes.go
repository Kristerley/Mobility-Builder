package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router{
    router := httprouter.New()

    router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

    // router.HandlerFunc(http.MethodPost, "/v1/exercises", app.createExerciseHandler)
    router.HandlerFunc(http.MethodGet, "/v1/exercises", app.showExerciseHandler)
    // router.HandlerFunc(http.MethodGet, "/v1/exercises/:id", app.showExerciseHandler)

    return router
}
