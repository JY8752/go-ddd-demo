package main

import (
	"database/sql"
	"ddd-demo/circle/appservice"
	"ddd-demo/circle/command"
	"ddd-demo/circle/domainservice"
	"ddd-demo/circle/repository"
	"ddd-demo/user"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "db/users.db")
	if err != nil {
		log.Fatal(err)
	}

	userRep := user.NewRepository(db)
	circleRep := repository.NewCircleRepository(db)
	cds := domainservice.NewCircleDomainService(circleRep)

	cas := appservice.NewCircleApplicationService(userRep, cds, circleRep)

	if err = cas.Create(command.CrateCircle{
		Id:   "userId",
		Name: "circleName",
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Complete!!")
}
