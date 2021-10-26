// Type declaration adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
// Gampangnya : membuat alias
// digunakan biasanya dengan tujuan agar lebih mudah dimengerti

package main

import "fmt"

func main() {
	type noKtp string
	type married bool

	var ktpDuma noKtp = "1231847817"
	var marriedStatusDuma = false
	fmt.Println("no ktp duma ", ktpDuma, "status ", marriedStatusDuma)

}
