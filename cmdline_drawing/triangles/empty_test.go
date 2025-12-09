package triangles

import "fmt"

func ExampleDrawEmptyTriangle() {
	DrawEmptyTriangle(1)
	fmt.Println("")
	DrawEmptyTriangle(2)
	fmt.Println("")
	DrawEmptyTriangle(3)
	fmt.Println("")
	DrawEmptyTriangle(4)
	fmt.Println("")
	DrawEmptyTriangle(6)
	fmt.Println("")
	DrawEmptyTriangle(8)
	fmt.Println("")
	DrawEmptyTriangle(13)
	// Output:
	// #
	// ##
	// # #
	// #  #
	// #   #
	// ######
}
