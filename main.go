package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const (
    host        = "localhost"
    port        = 5432
    user        = "postgres"
    password    = "not-a-real-password"
    dbname      = "kairos_test"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s passowrd=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)        
}