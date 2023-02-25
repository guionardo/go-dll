package main

import (
	"C"
)
import "fmt"

//export Entry
func Entry(name string) {
	fmt.Printf("%s OK", name)
}

func main() {
	fmt.Println("MAIN")
}
