package main

import (
	"encoding/json"
	"fmt"
)

type user2 struct {
	First string
	Age   int
}

func main() {
	u1 := user2{
		First: "James",
		Age:   32,
	}

	u2 := user2{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user2{
		First: "M",
		Age:   54,
	}

	users := []user2{u1, u2, u3}

	fmt.Println(users)

	sb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

}
