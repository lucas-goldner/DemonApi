package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "lucasgoldner"
	password = "lol123"
	dbname   = "demons"
)

//Demon Struct (Model)
type Demon struct {
	ID       string     `json:"id"`
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through demons and find with id
	for _, item := range demons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Demon{})
}

//Creates a new demon
func createDemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var demon Demon
	_ = json.NewDecoder(r.Body).Decode(&demon)
	demon.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not safe
	demons = append(demons, demon)
	json.NewEncoder(w).Encode(demon)
}

//Deletes a demon
func deleteDemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	for index, item := range demons {
		if item.ID == params["id"] {
			demons = append(demons[:index], demons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(demons)
}

//Updates a demon
func updateDemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	for index, item := range demons {
		if item.ID == params["id"] {
			demons = append(demons[:index], demons[index+1:]...)
			var demon Demon
			_ = json.NewDecoder(r.Body).Decode(&demon)
			demon.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not safe
			demons = append(demons, demon)
			json.NewEncoder(w).Encode(demon)
			return
		}
	}
	json.NewEncoder(w).Encode(demons)
}

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM demons")
	if err != nil {
		log.Fatal(err)
	}

	var demons []Demon

	for rows.Next() {
		var demon Demon
		rows.Scan(&demon.ID, &demon.Name, &demon.Weakness, &demon.Strength, &demon.Absorb, &demon.Imun, &demon.Reflect, &demon.Level)
		demons = append(demons, demon)
	}

	demonsBytes, _ := json.MarshalIndent(demons, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(demonsBytes)

	defer rows.Close()
	defer db.Close()
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var d Demon
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO demons (name, weakness, strength, absorb, imun, reflect, level) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(sqlStatement, d.Name, d.Weakness, d.Strength, d.Absorb, d.Imun, d.Reflect, d.Level)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data
	demons = append(demons, Demon{ID: "1", Level: 1, Name: "Jack Frost", Strength: "", Absorb: "Ice", Reflect: "", Weakness: "Fire", Imun: "", Attacks: []*Attacks{&Attacks{Name: "Bufu", Type: "Ice", Damage: 10, Description: "Light ice-attack"}}})
	demons = append(demons, Demon{ID: "2", Level: 10, Name: "Black Frost", Strength: "", Absorb: "Fire", Reflect: "Ice", Weakness: "", Imun: "", Attacks: []*Attacks{&Attacks{Name: "Bufu", Type: "Ice", Damage: 10, Description: "Light ice-attack"}, &Attacks{Name: "Agi", Type: "Fire", Damage: 10, Description: "Light fire-attack"}}})
	//Route Handlers / Endpoints
	r.HandleFunc("/api/demons", getDemons).Methods("GET")
	r.HandleFunc("/api/demons/{id}", getDemon).Methods("GET")
	r.HandleFunc("/api/demons", createDemon).Methods("POST")
	r.HandleFunc("/api/demons/{id}", deleteDemon).Methods("DELETE")
	r.HandleFunc("/api/demons/{id}", updateDemon).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4220", r))
}
