package practice

import "fmt"

func Variable() {
	fmt.Println("=== Variable ===")
	var x int
	var y float64
	var z string
	var name string

	var nameString = "nameString"

	name = "Hadi"
	fmt.Println("Integer:", x)
	fmt.Println("Float:", y)
	fmt.Println("String:", z)
	fmt.Println("Name:", name)
	fmt.Println("Name String without declare datatype:", nameString)

	name = "Luthfi"
	nameString = "update nameString"
	fmt.Println("Updated Name:", name)
	fmt.Println("Updated Name String without declare datatype:", nameString)

}
