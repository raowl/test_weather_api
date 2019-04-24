package handlers

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/raowl/test_weather_api/repos"
	"gopkg.in/mgo.v2"
	"net/http"
)

type AppContext struct {
	Db *mgo.Database
}

func (c *AppContext) WeathersHandler(w http.ResponseWriter, r *http.Request) {
	//params := context.Get(r, "params").(httprouter.Params)
	repo := repos.WeatherRepo{c.Db.C("weathers")}
	markers, err := repo.All()

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(markers)
}

func (c *AppContext) WeatherHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := repos.WeatherRepo{c.Db.C("weathers")}
	marker, err := repo.Find(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(marker)
}

func (c *AppContext) CreateWeatherHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*repos.WeatherResource)
	repo := repos.WeatherRepo{c.Db.C("weathers")}
	err := repo.Create(&body.Data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(body)
}
