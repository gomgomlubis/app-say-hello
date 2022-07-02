package main

import (
	"fmt"

	go_say_hello "github.com/gomgomlubis/go-say-hello"
)

func main() {
	fmt.Println(go_say_hello.SayHello())

	fmt.Println("3+3=", go_say_hello.Jumlah(3, 3))
}
