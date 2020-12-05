package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Demon Struct (Model)
type Demon struct {
	ID       int64      `json:"id"`
	Strength string     `json:"strength"`
	Weakness string     `json:"weakness"`
	Imun     string     `json:"imun"`
	Reflect  string     `json:"reflect"`
	Level    int64      `json:"level"`
	Attacks  []*Attacks `json:"attacks"`
}

type Attacks struct {
	Name        string `json:"name"`
	Damage      int64  `json:"damage"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

//Get All Demons
func getDemons(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Route Handlers / Endpoints

	log.Fatal(http.ListenAndServe(":4220", r))
}
