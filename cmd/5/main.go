package main

import (
	"context"
	"database/sql"
	"ddd-demo/domain/user"
	infra_user "ddd-demo/infrastructure/user"
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

	repository := infra_user.NewRepository(db)

	// テーブルを作成
	if db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	// Userをドメインモデルで表現
	name, err := user.NewFullName("yamanaka", "jun")
	if err != nil {
		log.Fatal(err)
	}

	u := user.Create(name)

	// ドメインサービスを作成
	service := user.NewService(repository)

	if service.Exists(u) {
		log.Fatal("already exist user " + u.String())
	}

	// ユーザーを永続化
	dto, err := repository.Create(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Complete created user %v", dto)
}
