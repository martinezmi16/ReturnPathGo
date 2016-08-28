package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)


//Creates animal type that will be encoded into JSON
type Animal struct {

    Name string `json:"Name"`
    LegCount int `json:"LegCount"`
    LifeSpan int `json:"LifeSpan"`
    IsEndangered bool `json:"IsEndangered"`

}

type Animals []Animal


//Main Page for Zoo
//This shows the main interface for interacting with the zoo website
func homePage( w http.ResponseWriter, r *http.Request ){

    fmt.Fprintf( w, "Welcome to the Zoo!" )
    fmt.Println( "Endpoint Hit: homePage" )
}


//Returns the animal indicated by the user
func returnAnimal( w http.ResponseWriter, r *http.Request ){

    fmt.Fprintf( w, "returns a specific animal" )
    fmt.Println( "Endpoint Hit: returnAnimal" )

}

//Returs all animals in the db
func returnAllAnimals( w http.ResponseWriter, r *http.Request ){

    animals := Animals{
        Animal{ Name: "cat", LegCount: 4, LifeSpan: 7, IsEndangered: false },
        Animal{ Name: "dog", LegCount: 4, LifeSpan: 10, IsEndangered: false },

    }

    fmt.Println( "Endpoint Hit: returnAllAnimals" )

    json.NewEncoder( w ).Encode( animals )

}

//Adds an animal to the db
func addAnimal( w http.ResponseWriter, r *http.Request ){

    fmt.Fprintf( w, "Adds animal to db" )
    fmt.Println( "Endpoint Hit: addAnimal" )

}

//Removes animal from db if it exists
func delAnimal( w http.ResponseWriter, r *http.Request ){

    fmt.Fprintf( w, "Removes animal from db" )
    fmt.Println( "Endpoint Hit: delAnimal" )

}


//Structure on how to handle requests
func handleRequests(){

    http.HandleFunc( "/", homePage )
    http.HandleFunc( "/all", returnAllAnimals )
    http.HandleFunc( "/single", returnAnimal )
    http.HandleFunc( "/delete", delAnimal )
    http.HandleFunc( "/add", addAnimal )
    log.Fatal( http.ListenAndServe( ":8080", nil ) )


}

func main() {

    handleRequests()

}





