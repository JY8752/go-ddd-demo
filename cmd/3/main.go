package main

import (
	"ddd-demo/user"
	"fmt"
	"log"
)

func main() {
	n, err := user.NewFullName("yamanaka", "junichi")
	if err != nil {
		log.Fatal(err)
	}

	u := user.New(n)
	uu := user.New(n)
	fmt.Println(u.Equals(uu)) // false 同じ名前のユーザーだけど識別子が異なる

	n, err = user.NewFullName("yamanaka", "jun")
	if err != nil {
		log.Fatal(err)
	}

	u.ChangeName(n) // ユーザーの名前は可変
	fmt.Println(n.String())
}
