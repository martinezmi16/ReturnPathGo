package main

import (
    "fmt"
    "log"
    "api"
    "net/http"

    "database/sql"
    _ "github.com/lib/pq"
    "api_db"
)






func main() {

    var err error
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", api_db.DB_USER, api_db.DB_PASSWORD, api_db.DB_NAME)
    api_db.DBCon, err = sql.Open( "postgres", dbinfo )

    if err == nil && api_db.DBCon != nil{

        router := api.NewRouter()

        log.Fatal( http.ListenAndServe( ":8080", router ) )
    } else {

            fmt.Println( "Access to DB denied" )

    }


}
