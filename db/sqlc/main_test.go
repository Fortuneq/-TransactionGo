package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:537j04222@localhost:5432/postgres?sslmode=disable"
)



var testQueries *Queries

func TestMain(m *testing.M){	
	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("cannot access database",err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}