package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler{
    router := httprouter.New()

    router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

    // router.HandlerFunc(http.MethodPost, "/v1/exercises", app.createExerciseHandler)
    router.HandlerFunc(http.MethodGet, "/v1/exercises/:purpose/:part", app.showExerciseHandler)
    // router.HandlerFunc(http.MethodGet, "/v1/exercises/:id", app.showExerciseHandler)
    return app.enableCORS(router)

}
