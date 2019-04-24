package main

import (
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"github.com/raowl/test_weather_api/handlers" //controllers
	"github.com/raowl/test_weather_api/repos"    //models
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	appC := handlers.AppContext{session.DB("test")}
	commonMiddleware := alice.New(context.ClearHandler, handlers.LoggingHandler, handlers.RecoverHandler, handlers.AcceptHandler)
	router := NewRouter()

	// user register
	router.Post("/api/v1/user", commonMiddleware.Append(handlers.ContentTypeHandler, handlers.BodyHandler(repos.UserResource{})).ThenFunc(appC.CreateUserHandler))
	// user authentication
	router.Post("/api/v1/user/auth", commonMiddleware.Append(handlers.ContentTypeHandler, handlers.BodyHandler(repos.UserResource{})).ThenFunc(appC.AuthUserHandler))
	router.Put("/api/v1/user", commonMiddleware.Append(handlers.ContentTypeHandler, handlers.AuthHandler, handlers.BodyHandler(repos.UserResource{})).ThenFunc(appC.UpdateUserHandler))
	router.Get("/api/v1/user/:id", commonMiddleware.Append(handlers.AuthHandler).ThenFunc(appC.UserHandler))

	// create weather
	router.Post("/api/v1/weather", commonMiddleware.Append(handlers.AuthHandler, handlers.ContentTypeHandler, handlers.BodyHandler(repos.WeatherResource{})).ThenFunc(appC.CreateWeatherHandler))
	router.Get("/api/v1/weather", commonMiddleware.Append(handlers.AuthHandler, handlers.ContentTypeHandler, handlers.BodyHandler(repos.WeatherResource{})).ThenFunc(appC.WeathersHandler))
	// get wheather by city
	router.Get("/api/v1/weather/:id", commonMiddleware.Append(handlers.AuthHandler).ThenFunc(appC.WeatherHandler))
	router.Get("/api/v1/weather/:id/history", commonMiddleware.ThenFunc(appC.WeatherHandler))
	//router.Get("/api/v1/user/username/:username", commonMiddleware.Append(handlers.AuthHandler).ThenFunc(appC.UserHandler))
	http.ListenAndServe(":8080", router)
}
