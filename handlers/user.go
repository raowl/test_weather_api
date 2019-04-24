package handlers

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/raowl/test_weather_api/config"
	"github.com/raowl/test_weather_api/repos"
	"gopkg.in/dgrijalva/jwt-go.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	//"path/filepath"
	"fmt"
	"time"
)

//POST: /api/v1/auth/ handler
func (c *AppContext) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var userregistered bool
	type Validationerror struct {
		Error_msg  string `bson:"error_msg,omitempty" json:"error_msg,omitempty"`
		Error_code string `bson:"error_code,omitempty" json:"error_code,omitempty"`
	}
	type ValidationerrorResource struct {
		Data Validationerror `json:"data"`
	}

	body := context.Get(r, "body").(*repos.UserResource)
	repo := repos.UserRepo{c.Db.C("users")}
	userregistered, err = repo.UserAlreadyExists(body.Data.Username)
	if err != nil {
		panic("err")
	}
	if userregistered {

		fmt.Println("el usuario ya existe")
		/* panic("error") */
		w.Header().Set("Content-Type", "application/vnd.api+json")
		w.WriteHeader(400)
		// e := ValidationerrorResource{Validationerror{"duplicate username", "125"}}
		e := Validationerror{"duplicate username", "125"}
		json.NewEncoder(w).Encode(e)
	} else {
		fmt.Println("el usuario no existe")
		err = repo.Create(&body.Data)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/vnd.api+json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(body)
	}
}

func (c *AppContext) UserHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := repos.UserRepo{c.Db.C("users")}

	var err error
	var user repos.UserResource

	//fmt.Printf("userId")
	//userId := context.Get(r, "userid").(string)

	if params.ByName("id") != "undefined" {
		fmt.Println("entro por aca")
		user, err = repo.Find(params.ByName("id"))
	}
	if err != nil {
		panic(err)
	}

	following, err := repo.GetFByIds(user.Data.Following)
	user.Data.FollowInfo = following
	fmt.Printf("Currently following00000000000000000000000000000000000000000000000000000000000000000000000000...\n")
	fmt.Printf("%+v\n", following)
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(user)
}

//POST: /api/v1/user/login/ handler
func (c *AppContext) AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	var (
		privateKey []byte
	)
	body := context.Get(r, "body").(*repos.UserResource)
	repo := repos.UserRepo{c.Db.C("users")}
	user_resource, err := repo.Authenticate(body.Data)
	if err != nil {
		panic(err)
	}
	//data := map[string]interface{"token": ""}
	//var data interface{}

	// Create JWT token
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["userid"] = user_resource.Data.Id
	// Expire in 5 mins
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	//token.Claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	//cprivateKey, _ := filepath.Abs(config.Info.PrivateKey)
	print("private key")
	print(config.Info.PrivateKey)
	privateKey, err = ioutil.ReadFile(config.Info.PrivateKey)
	if err != nil {
		panic(err)
	}
	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		panic(err)
	}

	fmt.Println("aca")
	//fmt.Println(user_resource.Data.Skills)
	data := map[string]interface{}{
		"token":      tokenString,
		"id":         user_resource.Data.Id,
		"skills":     user_resource.Data.Skills,
		"facebookid": user_resource.Data.FacebookId,
		"firstname":  user_resource.Data.FirstName,
		"lastname":   user_resource.Data.LastName,
		"following":  user_resource.Data.Following,
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(data)
}

func (c *AppContext) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("enter to updateuser handler")
	//params := context.Get(r, "params").(httprouter.Params)
	body := context.Get(r, "body").(*repos.UserResource)
	//fmt.Printf("userId")
	userId := context.Get(r, "userid").(string)
	body.Data.Id = bson.ObjectIdHex(userId)
	//fmt.Printf("BODY DATA")
	//fmt.Printf("%+v\n", body.Data)
	//body.Data.Skills = []bson.ObjectId{bson.ObjectIdHex(params.ByName("id"))}
	//body.Data.Following = followingIdArr
	//body.Data.Skills = []repos.Skill{{userId}}
	repo := repos.UserRepo{c.Db.C("users")}
	err := repo.Update(&body.Data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(204)
	w.Write([]byte("\n"))
}
