package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/test_weather_api/config"
	"github.com/test_weather_api/utils"
	"gopkg.in/dgrijalva/jwt-go.v2"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

var (
	PublicKey []byte
)

func init() {
	PublicKey, _ = ioutil.ReadFile(config.Info.PublicKey)
}

func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				utils.WriteError(w, utils.ErrInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

//func (c *appContext) authHandler(next http.Handler) http.Handler {
func AuthHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.ParseFromRequest(r, func(t *jwt.Token) (interface{}, error) {
			return PublicKey, nil
		})

		if err != nil {
			utils.WriteError(w, utils.ErrInternalServer)
			return
		}
		if token.Valid == false {
			fmt.Printf("token malo")
			//YAY!
			utils.WriteError(w, utils.ErrInternalServer)
			return
		}

		fmt.Printf("token valido\n")

		context.Set(r, "userid", token.Claims["userid"].(string))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func AcceptHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") != "application/vnd.api+json" {
			utils.WriteError(w, utils.ErrNotAcceptable)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func ContentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/vnd.api+json" {
			utils.WriteError(w, utils.ErrUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func BodyHandler(v interface{}) func(http.Handler) http.Handler {
	fmt.Printf("Inside BodyHandler\n")
	t := reflect.TypeOf(v)

	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)

			if err != nil {
				utils.WriteError(w, utils.ErrBadRequest)
				return
			}

			if next != nil {
				context.Set(r, "body", val)
				next.ServeHTTP(w, r)
			}
		}

		return http.HandlerFunc(fn)
	}

	return m
}
