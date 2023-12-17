package main

import (
	"database/sql"
	application_circle "ddd-demo/application/circle"
	domain_circle "ddd-demo/domain/circle"
	"ddd-demo/infrastructure/circle"
	"ddd-demo/infrastructure/user"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "db/users.db")
	if err != nil {
		log.Fatal(err)
	}

	userRep := user.NewRepository(db)
	circleRep := circle.NewCircleRepository(db)
	cds := domain_circle.NewDomainService(circleRep)

	cas := application_circle.NewApplicationService(userRep, cds, circleRep)

	if err = cas.Create(application_circle.CrateCommand{
		Id:   "userId",
		Name: "circleName",
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Complete!!")
}
