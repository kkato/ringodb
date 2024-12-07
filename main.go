package main

import (
	"fmt"
	"os"

	"github.com/kkato/ringodb/backend"
)

func main() {
	data := []byte("Hello, World!")
	err := backend.SaveData1("hello.txt", data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Data saved to hello.txt")
}
