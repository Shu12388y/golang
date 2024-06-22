package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Hello this working")
	presentTime := time.Now()
	fmt.Print(presentTime.Format("01-02-2006 15.04.05 Monday"))              // for formating time you have to use the parameter given in Format()
}
