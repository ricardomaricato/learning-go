package main

import "fmt"

/*
- Partindo do código abaixo, utilize marshal para transformar []user2 em JSON.
*/

type user struct {
	first string
	age   int
}

func main() {
	u1 := user{
		first: "James",
		age:   32,
	}

	u2 := user{
		first: "Moneypenny",
		age:   27,
	}

	u3 := user{
		first: "M",
		age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
}
