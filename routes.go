package main

import ( 

    "net/http"

)


type Route struct {

    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc

}

type Routes []Route



var routes = Routes{

    Route{
        "homePage",
        "GET",
        "/",
        homePage,
    },

    Route{
        "returnAllAnimals",
        "GET",
        "/animals",
        returnAllAnimals,
    },

    Route{
        "returnOneAnimal",
        "GET",
        "/animals/{id}",
        returnOneAnimal,
    },

    Route{
        "addAnimal",
        "POST",
        "/animals",
        addAnimal,
    },
    

}








