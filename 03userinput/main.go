package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello this is user input"
	fmt.Println(welcome)
	fmt.Print("enter your rating ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Print("The rating is ", input)

}
