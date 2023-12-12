package main

import (
	"context"
	"database/sql"
	"ddd-demo/user"
	"ddd-demo/user/application"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "db/users.db")
	if err != nil {
		log.Fatal(err)
	}

	r := user.NewRepository(db)
	s := user.NewService(r)
	as := application.NewService(r, s)

	result, err := as.Create(ctx, "yamanaka", "junichi")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Complete!! user: %v", result)
}
