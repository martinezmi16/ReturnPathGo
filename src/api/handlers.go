package api

import (

    "fmt"
    "log"
    "io"
    "io/ioutil"

    "net/http"
"html/template"
    "encoding/json"

    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/lib/pq"
    "api_db"


    "strconv"
)




//Main Page for Zoo
//This shows the main interface for interacting with the zoo website
func homePage( w http.ResponseWriter, r *http.Request ){

    //Indicates endpoint has been hit
    fmt.Fprintf( w, "Welcome to the Zoo!" )
    fmt.Println( "Endpoint Hit: homePage" )

    //Creates template and executes html


}



//Returns a single animal based on its id
func returnOneAnimal( w http.ResponseWriter, r *http.Request ) {

    //Gets the ID from the arguments
    vars := mux.Vars( r )
    id := vars[ "id" ]
    id_int, err := strconv.Atoi( id )

    if err != nil{
        fmt.Fprintf( w, "Please enter valid animal ID" )

    } else {

    

        //Creates new animal struct
        currentAnimal := new( Animal )

        //Queries the database for animal with requested ID
        err = api_db.DBCon.QueryRow("SELECT * FROM animals WHERE id=$1", id_int).
                        Scan( &currentAnimal.Id, &currentAnimal.Name,
                                &currentAnimal.LegCount, &currentAnimal.LifeSpan,
                                &currentAnimal.IsEndangered, &currentAnimal.CreatedTime,
                                &currentAnimal.UpdatedTime  )

        //Error handling
        switch{
        case err == sql.ErrNoRows:
            log.Printf("No animals with that ID")
        case err != nil:
            log.Fatal(err)
        default:
            fmt.Printf( "%s is the name", currentAnimal.Name )
        }


        //Converts animal struct to json object
        json.NewEncoder( w ).Encode( currentAnimal )

    }
}

//Returs all animals in the db
func returnAllAnimals( w http.ResponseWriter, r *http.Request ){


    var animals []*Animal

    //Indicates that the endpoint has been hit
    fmt.Println( "Endpoint Hit: returnAllAnimals" )

    //Queries database to find all animals
    rows, err := api_db.DBCon.Query("SELECT * FROM animals")
    checkErr( err )

    
    //Iterates through returned rows
    for rows.Next() {

        //Reads fields into animal object and appends to animal list
        currentAnimal := new( Animal)
        err = rows.Scan( &currentAnimal.Id, &currentAnimal.Name, &currentAnimal.LegCount,
                           &currentAnimal.LifeSpan, &currentAnimal.IsEndangered,
                                &currentAnimal.CreatedTime, &currentAnimal.UpdatedTime  )
        checkErr( err )
        fmt.Printf( "%s is the name", currentAnimal.Name )
        animals = append( animals, currentAnimal )

    }



    //Indicates status and converts object to json
    w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
    w.WriteHeader( http.StatusOK )

    json.NewEncoder( w ).Encode( animals )

}

//Adds an animal to the db
func addAnimal( w http.ResponseWriter, r *http.Request ){

    //Indicates endpoint has been hit
    fmt.Fprintf( w, "Adds animal to db" )
    fmt.Println( "Endpoint Hit: addAnimal" )


    //Creates new animal to be read into
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

    //Inserts animal to the database based on information given to the server
    var newID int
    err = api_db.DBCon.QueryRow("INSERT INTO animals(name, leg_count, lifespan, is_endangered, created_at, updated_at) VALUES ( $1, $2, $3, $4, NOW(), NOW() ) Returning id", newAnimal.Name, newAnimal.LegCount, newAnimal.LifeSpan, newAnimal.IsEndangered).Scan( &newID )
        switch{
        case err == sql.ErrNoRows:
            log.Printf("No animals with that ID")
        case err != nil:
            log.Fatal(err)

        }


    //Returns new animal in json format
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)

    if err := json.NewEncoder(w).Encode( newAnimal ); err != nil {
        panic(err)
    }



}

//Removes animal from db if it exists
func delAnimal( w http.ResponseWriter, r *http.Request ){

    //Indicates endpoint has been hit
    fmt.Fprintf( w, "Removes animal from db" )
    fmt.Println( "Endpoint Hit: delAnimal" )

    vars := mux.Vars( r )
    id := vars[ "id" ]
    id_int, err := strconv.Atoi( id )

    if err != nil && id_int > 0{
        fmt.Fprintf( w, "Please enter valid animal ID" )

    } else {

    



        stmt, err := api_db.DBCon.Prepare("DELETE FROM animals WHERE id=$1")
        checkErr( err )



        res, err := stmt.Exec(1)
        checkErr(err)

        affect, err := res.RowsAffected()
        checkErr(err)

        fmt.Println(affect, "rows changed")



        fmt.Fprintf( w, "ID: " + id )
        json.NewEncoder( w ).Encode( affect )

    }

}

//Function for checking for errors
func checkErr(err error) {
    if err != nil {
        panic( err )
    }
}
