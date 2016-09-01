package main

import (
    "fmt"
    "log"
    "net/http"

    "database/sql"
    _ "github.com/lib/pq"
)


const (

    DB_USER = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME = "return_path_api_development"

)

var DBCon *sql.DB



func main() {

    var err error
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    DBCon, err = sql.Open( "postgres", dbinfo )

    if err == nil && DBCon != nil{

        router := NewRouter()

        log.Fatal( http.ListenAndServe( ":8080", router ) )
    } else {

            fmt.Println( "Access to DB denied" )

    }


}
