package main

import (
	"io"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"

)

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}

func DataIndex(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Query Home")
}

func DataShow(rw http.ResponseWriter, r *http.Request) {

    var book Book
    var books []Book
    var title, author string

    db, err := sql.Open("sqlite3", "books.sqlite")

    if err != nil {
        panic(err)
    }

    vars := mux.Vars(r)
    queryString := vars["queryString"]




    sql := fmt.Sprintf("select * from books where title like '%s%s%s'","%", queryString, "%")

    rows, err := db.Query(sql, queryString)

    if err != nil {
        panic(err)
    }

    for rows.Next(){

        rows.Scan(&title, &author)
        book.Author = author
        book.Title = title

        books = append(books, book)

        

    }

    js, err := json.Marshal(books)
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    rw.Header().Set("Content-Type", "application/json")
    rw.Write(js)

    db.Close()
}

func DataInsert(rw http.ResponseWriter, r *http.Request) {
    var book Book

    db, err := sql.Open("sqlite3", "books.sqlite")
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &book); err != nil {
        rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
        rw.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(rw).Encode(err); err != nil {
            panic(err)
        }
    }

    fmt.Printf("%+v\n", book)
    db.Exec("INSERT INTO books VALUES(?, ?)", book.Title, book.Author)

    rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
    rw.WriteHeader(http.StatusCreated)
    
    db.Close()

}