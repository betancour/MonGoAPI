package book

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "books"

// DOCNAME the name of the document
const DOCNAME = "book"

// GetBooks returns the list of Books
func (r Repository) GetBooks() Books {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Books{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddBook inserts a Book in the DB
func (r Repository) AddBook(book Book) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	book.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(book)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateBook updates a Book in the DB
func (r Repository) UpdateBook(book Book) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	book.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(book.ID, book)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteBook deletes a Book
func (r Repository) DeleteBook(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "404"
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove book
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "500"
	}

	// Write status
	return "200"
}
