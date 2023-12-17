package main

import (
	"context"
	"database/sql"
	app_user "ddd-demo/application/user"
	"ddd-demo/domain/user"
	infra_user "ddd-demo/infrastructure/user"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "db/users.db")
	if err != nil {
		log.Fatal(err)
	}

	r := infra_user.NewRepository(db)
	s := user.NewService(r)
	as := app_user.NewApplicationService(r, s)

	result, err := as.Create(ctx, "yamanaka", "junichi")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Complete!! user: %v", result)
}
