package repos

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Weather struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"name,omitempty"`
	time     string        `json:"author,omitempty"`
}

type WeatherCollection struct {
	Data []Weather `json:"data"`
}

type WeatherResource struct {
	Data Weather `json:"data"`
}

type WeatherRepo struct {
	Coll *mgo.Collection
}

func (r *WeatherRepo) All() (WeatherCollection, error) {

	result := WeatherCollection{[]Weather{}}
	err := r.Coll.Find(nil).All(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *WeatherRepo) Find(id string) (WeatherResource, error) {
	result := WeatherResource{}
	err := r.Coll.FindId(bson.ObjectIdHex(id)).One(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *WeatherRepo) Create(weather *Weather) error {
	id := bson.NewObjectId()
	_, err := r.Coll.UpsertId(id, weather)
	if err != nil {
		return err
	}

	weather.Id = id

	return nil
}

func (r *WeatherRepo) Delete(id string) error {
	err := r.Coll.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return nil
}
