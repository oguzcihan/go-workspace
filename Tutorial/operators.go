package main

import "fmt"

/*
	+ => Adds two operands
	- => Subtracts two operands
	* => Multiplies two operands
	/ => Divides two operands
	% => Returns the remainder of two operands
	& => Bitwise AND of two operands
	| => Bitwise OR of two operands
	^ => Bitwise XOR of two operands
	% => Mod
	++ => Increments the operand by 1
	-- => Decrements the operand by 1
*/
func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	area += LENGTH
	fmt.Printf("Area is %d", area)
}
