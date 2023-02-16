package controllers


import (
	"fmt"
	"main.go/models"
	"github.com/jullienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"net/http"
)



type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		res.WriteHeader(http.StatusNotFound)	
	}
	
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if uc.session.DB("the new mongodb intergration with golang").C("users").FindId(oid).One(&u); err  != nil{
		res.WriteHeader(404)
		return
	}

	uj ,err := json.Marshal(u)
	if err != nil{
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	fmt.Fprintf(res, "%s\n", uj)	
}


func (uc UserController) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	u := models.User{}

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil{
		panic(err)
	}

	u.Id = bson.NewObjectId()

	uc.session.DB("the new mongodb integration with golang").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil{
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(201)
	fmt.Fprintf(res, "%s\n", uj)	
}




func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		res.WriteHeader(http.StatusNotFound)	
	}
	
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if uc.session.DB("the mongodb integration with golang").C("users").RemoveId(oid); err != nil{
		res.WriteHeader(404)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200 )
	fmt.Fprintf(res, "%s\n", u)	
}


