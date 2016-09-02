package api_db

import (

    "database/sql"
    _ "github.com/lib/pq"

)

const (

    DB_USER = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME = "return_path_api_development"

)

var DBCon *sql.DB
