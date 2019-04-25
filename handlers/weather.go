package handlers

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/raowl/test_weather_api/repos"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type AppContext struct {
	Db *mgo.Database
}

//get all
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

// get by id

func (c *AppContext) WeatherHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := repos.WeatherRepo{c.Db.C("weathers")}
	weather, err := repo.Find(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	// add viewedBy
	// TODO: this out of time dont seem to work, commented for compiling need to check again about it in golang
	//userId := bson.ObjectIdHex(context.Get(r, "userid").(string))
	//body := context.Get(r, "body").(*repos.WeatherResource)
	//weather.ViewedBys = []repo.ViewedBys{{userId}}
	//err := repo.Update(weather)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(weather)
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
