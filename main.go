package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const (
    host        = "localhost"
    port        = 5432
    user        = "kairos_test_rw"
    password    = "not-a-real-password"
    dbname      = "kairos_test"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s passowrd=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)        
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully Connected!")
}