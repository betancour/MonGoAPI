package book

import "gopkg.in/mgo.v2/bson"

//Books represents a single book
type Books struct {
	Book struct {
		Author  string `json:"author"`
		Current struct {
			Edition  time.Time `json:"edition"`
			Location string    `json:"location"`
			Owner    string    `json:"owner"`
		} `json:"current"`
		Details struct {
			Categories []string `json:"categories"`
			Isbn       string   `json:"isbn"`
			Language   string   `json:"language"`
			Pages      int      `json:"pages"`
			Publisher  string   `json:"publisher"`
			Year       int      `json:"year"`
		} `json:"details"`
		Genre  string `json:"genre"`
		Status struct {
			Donated bool `json:"donated"`
			Loan    bool `json:"loan"`
			Lost    bool `json:"lost"`
			Sold    bool `json:"sold"`
		} `json:"status"`
		Title string `json:"title"`
	} `json:"book"`
}

//Books is an array of book
type Books []Book
