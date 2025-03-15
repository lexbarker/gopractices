package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Get_Name()
}

func Get_Name() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("hello world, whats your name?")
	input, _ := reader.ReadString('\n')

	fmt.Printf("Hi %s}", input)
}
