package main

import (

    "fmt"
    "log"
    "io"
    "io/ioutil"

    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/lib/pq"


    "strconv"
)



//Function for checking for errors
func checkErr(err error) {
    if err != nil {
        panic( err )
    }
}


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


//Returns a single animal based on its id
func returnOneAnimal( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    id := vars[ "id" ]
    id_int, err := strconv.Atoi( id )

    if err != nil{
        fmt.Fprintf( w, "Please enter valid animal ID" )

    } else {

    


        currentAnimal := new( Animal)
        err = DBCon.QueryRow("SELECT * FROM animals WHERE id=$1", id_int).
                        Scan( &currentAnimal.Id, &currentAnimal.Name,
                                &currentAnimal.LegCount, &currentAnimal.LifeSpan,
                                &currentAnimal.IsEndangered, &currentAnimal.CreatedTime,
                                &currentAnimal.UpdatedTime  )
        switch{
        case err == sql.ErrNoRows:
            log.Printf("No animals with that ID")
        case err != nil:
            log.Fatal(err)
        default:
            fmt.Printf( "%s is the name", currentAnimal.Name )
        }




        fmt.Fprintf( w, "ID: " + id )
        json.NewEncoder( w ).Encode( currentAnimal )

    }
}

//Returs all animals in the db
func returnAllAnimals( w http.ResponseWriter, r *http.Request ){

    var animals []*Animal

    fmt.Println( "Endpoint Hit: returnAllAnimals" )

    rows, err := DBCon.Query("SELECT * FROM animals")
    checkErr( err )

    

    for rows.Next() {
        currentAnimal := new( Animal)
        err = rows.Scan( &currentAnimal.Id, &currentAnimal.Name, &currentAnimal.LegCount,
                           &currentAnimal.LifeSpan, &currentAnimal.IsEndangered,
                                &currentAnimal.CreatedTime, &currentAnimal.UpdatedTime  )
        checkErr( err )
        fmt.Printf( "%s is the name", currentAnimal.Name )
        animals = append( animals, currentAnimal )

    }




    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    w.WriteHeader( http.StatusOK )

    json.NewEncoder( w ).Encode( animals )

}

//Adds an animal to the db
func addAnimal( w http.ResponseWriter, r *http.Request ){

    fmt.Fprintf( w, "Adds animal to db" )
    fmt.Println( "Endpoint Hit: addAnimal" )



    var newAnimal Animal
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

    if err != nil {
        panic(err)
    }

    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    if err := json.Unmarshal(body, &newAnimal); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    var newID int
    err = DBCon.QueryRow("INSERT INTO animals(name, leg_count, lifespan, is_endangered, created_at, updated_at) VALUES ( $1, $2, $3, $4, NOW(), NOW() ) Returning id", newAnimal.Name, newAnimal.LegCount, newAnimal.LifeSpan, newAnimal.IsEndangered).Scan( &newID )
        switch{
        case err == sql.ErrNoRows:
            log.Printf("No animals with that ID")
        case err != nil:
            log.Fatal(err)

        }



    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)

    if err := json.NewEncoder(w).Encode( newAnimal ); err != nil {
        panic(err)
    }



}

//Removes animal from db if it exists
func delAnimal( w http.ResponseWriter, r *http.Request ){

    fmt.Fprintf( w, "Removes animal from db" )
    fmt.Println( "Endpoint Hit: delAnimal" )

}
