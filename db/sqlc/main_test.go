package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var TestQueries *Queries

func TestMain(m *testing.M){
	//var err error
	conn,err:=pgx.Connect(context.Background(),"postgresql://user1:test@localhost/hospital?sslmode=disable")
	if err!=nil{
		log.Fatal("Couldn't connect to db")
	}
	TestQueries=New(conn)
	os.Exit(m.Run())
}