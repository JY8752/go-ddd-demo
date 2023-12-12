package main

import (
	"context"
	"database/sql"
	"ddd-demo/user"
	"fmt"
	"log"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed sql/schema.sql
var ddl string

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "db/users.db")
	if err != nil {
		log.Fatal(err)
	}

	repository := user.NewRepository(db)

	// テーブルを作成
	if err = repository.CreateTables(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	// Userをドメインモデルで表現
	name, err := user.NewFullName("yamanaka", "jun")
	if err != nil {
		log.Fatal(err)
	}

	u := user.NewFromName(name)

	// ドメインサービスを作成
	service := user.NewService(repository)

	if service.Exists(u) {
		log.Fatal("already exist user " + u.String())
	}

	// ユーザーを永続化
	dto, err := repository.CreateUser(ctx, u.Id().String(), u.Name().FirstName(), u.Name().LastName())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Complete created user %v", dto)
}
