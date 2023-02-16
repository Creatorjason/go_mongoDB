package main

import (
	"log"
	"net/http"
	"main.go/controllers"
	"github.com/jullienschmidt/httprouter"
	"gopkg.in/mgo.v2" // MongoDB connector
)


func main(){
	router := httprouter.New()
	userSession := controllers.NewUserController(getSession())
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)


	log.Fatal(http.ListenAndServe(":9000", router))
}


func getSession() *mgo.Session{
	session, err := mgo.Dial("mongodb://localhost:27107") // establishes a database connection
	if err != nil{
		panic(err)
	}

	return session

	
}