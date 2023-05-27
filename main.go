package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Note struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Fake DB
var notes []Note

// middleware, helper - file
func (n Note) IsEmpty() bool {
	return n.Title == "" && n.Description == ""
}

func main() {
	fmt.Println("API - Notes App")
	r := mux.NewRouter()

	// seeding data
	notes = append(notes, Note{Id: "1", Title: "Title1", Description: "Description1", User: &User{Name: "User1", Email: "User1@gmail.com"}})
	notes = append(notes, Note{Id: "2", Title: "Title2", Description: "Description2", User: &User{Name: "User2", Email: "User2@gmail.com"}})
	notes = append(notes, Note{Id: "3", Title: "Title3", Description: "Description3", User: &User{Name: "User3", Email: "User3@gmail.com"}})
	notes = append(notes, Note{Id: "4", Title: "Title4", Description: "Description4", User: &User{Name: "User4", Email: "User4@gmail.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/notes", getAllNotes).Methods("GET")
	r.HandleFunc("/note/{id}", getNoteById).Methods("GET")
	r.HandleFunc("/note", createNote).Methods("POST")
	r.HandleFunc("/note/{id}", updateNote).Methods("PUT")
	r.HandleFunc("/note/{id}", deleteNoteById).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":3000", r))
	http.Handle("/", r)
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Notes API</h1>"))
}

func getAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all notes")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func getNoteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Note")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through notes, find matching id and return the response
	for _, note := range notes {
		if note.Id == params["id"] {
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	json.NewEncoder(w).Encode("No Note found with given id")
	return
}

func createNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Note")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)

	if note.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside")
		return
	}

	// generate unique id and convert to string
	rand.Seed(time.Now().UnixNano())
	note.Id = strconv.Itoa(rand.Intn(100))
	notes = append(notes, note)
	json.NewEncoder(w).Encode(note)
	return
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update the Note")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request

	params := mux.Vars(r)
	var id string = params["id"]

	// loop, id, remove, add with my ID
	for index, note := range notes {
		if note.Id == id {
			notes = append(notes[:index], notes[index+1:]...)
			var updated_note Note
			_ = json.NewDecoder(r.Body).Decode(&updated_note)
			note.Id = id
			notes = append(notes, updated_note)
			json.NewEncoder(w).Encode(updated_note)
			return 
		}
	}
	//TODO: send a response when id is not found
}

func deleteNoteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Note")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var id string = params["id"]

	for index, note := range notes {
		if note.Id == id {
			notes = append(notes[:index], notes[index+1:]...)
			json.NewEncoder(w).Encode(note)
			return
		}
	}
}
