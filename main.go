package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Demon Struct (Model)
type Demon struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Strength string     `json:"strength"`
	Weakness string     `json:"weakness"`
	Imun     string     `json:"imun"`
	Absorb   string     `json:"absorb"`
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

// Init books var as a slice Book struct
var demons []Demon

//Get All Demons
func getDemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(demons)
}

//Get One Demon
func getDemon(w http.ResponseWriter, r *http.Request) {

}

//Creates a moveset
func createDemon(w http.ResponseWriter, r *http.Request) {

}

//Deletes a moveset
func deleteDemon(w http.ResponseWriter, r *http.Request) {

}

//Updates a moveset
func updateDemon(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data
	demons = append(demons, Demon{ID: 1, Level: 1, Name: "Jack Frost", Strength: "", Absorb: "Ice", Reflect: "", Weakness: "Fire", Imun: "", Attacks: []*Attacks{&Attacks{Name: "Bufu", Type: "Ice", Damage: 10, Description: "Light ice-attack"}}})
	demons = append(demons, Demon{ID: 2, Level: 10, Name: "Black Frost", Strength: "", Absorb: "Fire", Reflect: "Ice", Weakness: "", Imun: "", Attacks: []*Attacks{&Attacks{Name: "Bufu", Type: "Ice", Damage: 10, Description: "Light ice-attack"}, &Attacks{Name: "Agi", Type: "Fire", Damage: 10, Description: "Light fire-attack"}}})
	//Route Handlers / Endpoints
	r.HandleFunc("/api/demons", getDemons).Methods("GET")
	r.HandleFunc("/api/demons/{id}", getDemon).Methods("GET")
	r.HandleFunc("/api/demons", createDemon).Methods("POST")
	r.HandleFunc("/api/demons/{id}", deleteDemon).Methods("DELETE")
	r.HandleFunc("/api/demons/{id}", updateDemon).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4220", r))
}
