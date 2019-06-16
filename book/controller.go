package book

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	Repository Repository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	books := c.Repository.GetBooks() // list of all books
	log.Println(books)
	data, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddBook POST /
func (c *Controller) AddBook(w http.ResponseWriter, r *http.Request) {

	var book Book
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error AddBook", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddBook", err)
	}
	if err := json.Unmarshal(body, &book); err != nil { // unmarshall body contents as a type Book
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddBook unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.AddBook(book) // adds the book to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateBook UPDATE /
func (c *Controller) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error UpdateBook", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddaUpdateBook", err)
	}
	if err := json.Unmarshal(body, &book); err != nil { // unmarshall body contents as a type Book
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateBook unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.Repository.UpdateBook(book) // adds the book to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteBook DELETE /
func (c *Controller) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] // param id

	if err := c.Repository.DeleteBook(id); err != "" { // delete a book by id
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	return
}
