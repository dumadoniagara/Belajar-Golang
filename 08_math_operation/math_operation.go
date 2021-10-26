package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println("hasilnya adalah ", result)

	// Berlaku juga augmented Assignment
	var a = 10
	a += 15
	fmt.Println("hasilnya a adalah ", a)

	// Berlaku juga unary operator
	var b = 11
	b++
	fmt.Println("hasilnya b adalah ", b)

}
