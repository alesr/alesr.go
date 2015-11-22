package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/alesr/alesr.go/rest/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// instantiate a new router
	r := httprouter.New()

	// get a BarController instance
	barCtrl := controllers.NewBarController(getSession())

	// get a bar resource
	r.GET("/bar/:id", barCtrl.GetBar)

	r.POST("/bar", barCtrl.CreateBar)

	r.DELETE("/bar/:id", barCtrl.RemoveBar)

	// fire up the server
	http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session {

	// connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// check if connecrion error, is mongo running?
	if err != nil {
		log.Fatal("mongo error", err)
	}

	return s
}

// curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Harmonia", "description": "best bar evora", "latitude": -80, "longitude": 20}' http://localhost:3000/bar
// curl http://localhost:3000/user/5497246c380a967ff1000003
