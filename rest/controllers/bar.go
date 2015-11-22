package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alesr/alesr.go/rest/models"
	"github.com/julienschmidt/httprouter"
)

type (
	// BarController representes the controller for operating the Bar resource
	BarController struct {
		session *mgo.Session
	}
)

func NewBarController(s *mgo.Session) *BarController {
	return &BarController{s}
}

// GetBat retrieves a single bar resource
func (barCtrl BarController) GetBar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// grab id
	id := p.ByName("id")

	// verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// grab ObjectId
	oid := bson.ObjectIdHex(id)

	// stub bar
	b := models.Bar{}

	// fetch bar
	if err := barCtrl.session.DB("kON_DB").C("bars").FindId(oid).One(&b); err != nil {
		w.WriteHeader(404)
		return
	}

	// marshal provided interface into JSON structure
	bj, err := json.Marshal(b)
	if err != nil {
		log.Fatal("error on marshling data into JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", bj)
}

// CreateBar creates a new bar resource
func (barCtrl BarController) CreateBar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// stub a bar to be populated from the body
	b := models.Bar{}

	// populate the bar data
	json.NewDecoder(r.Body).Decode(&b)

	// add an ID
	b.ID = bson.NewObjectId()

	// write the bar to mongo
	barCtrl.session.DB("kON_DB").C("bars").Insert(b)

	// marshal provided interface into JSON structure
	bj, err := json.Marshal(b)
	if err != nil {
		log.Fatal("error on marshaling created bar data to JSON", err)
	}

	// write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", bj)
}

// RemoveBar removes an existing bar resourse
func (barCtrl BarController) RemoveBar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// grab id
	id := p.ByName("id")

	// verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// grab ObjectId
	oid := bson.ObjectIdHex(id)

	// remove user
	if err := barCtrl.session.DB("kON_DB").C("bars").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// write status
	w.WriteHeader(200)
}
