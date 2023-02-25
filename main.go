package main

import (
	"C"
)
import "fmt"

//export Entry
func Entry() {
	fmt.Println("OK")
}

func main() {
	fmt.Println("MAIN")
}
