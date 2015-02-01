package main

import ( 
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
    Title  string `json:"title"`
    Author string `json:"author"`
}

func main() {

	//db := NewDB()

	//InsertBooks(db)
	r := NewRouter()

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
 

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "books.sqlite")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists books(title text, author text)")
	if err != nil {
		panic(err)
	}

	return db
}
