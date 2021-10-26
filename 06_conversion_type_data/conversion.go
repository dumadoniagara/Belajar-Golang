package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32) // type data 32 akan langsung dikonversi menjadi tipe data 64

	// cuma harus hati2 jika dikonversi ke yang lebih kecil karena ga cukup panjangnya.

	fmt.Println("nilai 32 =", nilai32, "nilai64 =", nilai64)

	// konversi byte ke string

	var d = "Duma Doniagara Sambora"[2]

	var dString = string(d)

	fmt.Println("nilai yang diambil dari d adalah ", dString)

}
