package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/islamghany/go-auth/api"
	_ "github.com/lib/pq"
)

func main() {

	conn, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Database connection has been established successfully.")

	fmt.Println("server is initalized on port 4000")

	server := api.NewServer("islam")
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
