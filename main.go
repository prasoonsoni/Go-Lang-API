package main

import (
	"encoding/json"
	"fmt"
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
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Fake DB
var notes []Note

// middleware, helper - file
func (n Note) IsEmpty() bool {
	return n.Title == "" && n.Description == ""
}

func main() {

}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Notes API</h1>"))
}

func getAllNotes(w http.ResponseWriter, r *http.Response) {
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
