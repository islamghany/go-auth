package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"log"

	"github.com/islamghany/go-auth/api"
	db "github.com/islamghany/go-auth/db/sqlc"
	_ "github.com/lib/pq"
)

// Your secret key should be a cryptographically secure random string with an underlying entropy of at least 32 bytes (256 bits).
const JWT_SECRET = "pei3einoh0Beem6uM6Ungohn2heiv5lah1ael4joopie5JaigeikoozaoTew2Eh6"

func main() {

	conn, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Database connection has been established successfully.")

	fmt.Println("server is initalized on port 4000")

	store := db.New(conn)
	server := api.NewServer(store)

	log.Fatal(server.Start(4000))
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:secret@localhost:5431/auth?sslmode=disable")

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
