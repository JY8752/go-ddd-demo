package main

import (
	"ddd-demo/money"
	"ddd-demo/user"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

func main() {
	n, err := user.NewFullName("yamanaka", "junichi")
	if err != nil {
		log.Fatal(err)
	}

	n2, err := user.NewFullName("yamanaka", "junichi")
	if err != nil {
		log.Fatal(err)
	}

	// _, err = user.NewFullName("", "junichi")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(n == n2)
	fmt.Println(n.FirstName(), n.LastName(), n.FullName())
	fmt.Println(n.Equals(n2))
	fmt.Println(n)

	m1 := money.NewMoney(decimal.NewFromFloat(1.2), money.JPY)
	m2 := money.NewMoney(decimal.NewFromFloat(1.2), money.JPY)
	m3 := money.NewMoney(decimal.NewFromFloat(3.2), money.USD)

	fmt.Println(m1, m2, m3)
	fmt.Println(m1.Equals(m3))
}
