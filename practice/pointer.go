package practice

import "fmt"

func Pointer() {
	fmt.Println("=== Pointer ===")
	name := "Luthfi"
	namePointer := &name
	fmt.Println("Memory Address of Name:", namePointer)
	fmt.Println("Name:", *namePointer)
}