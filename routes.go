package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "DataIndex",
        "GET",
        "/database",
        DataIndex,
    },
    Route{
        "DataShow",
        "GET",
        "/database/{bookTitle}",
        DataShow,
    },
    Route{
        "DataInsert",
        "POST",
        "/database",
        DataInsert,
    },
}