package main

import "fmt"

func main() {
	// ketika mendeklarasikan constant WAJIB mendeklarasikan isi dari konstanta nya
	// sekali const di deklarasikan tidak akan bisa di ubah lagi

	const firstName = "Duma"
	const age = 25
	const space = " " //jika dia mendeklarasikan constant tapi ga langsung dipake ga akan error beda dengan variable

	// Kalo konstan ga bisa mendeklarasikan dengan := (titik dua sama dengan)
	// tapi bisa juga dengan menggunakan multiple declaration

	const (
		friendName = "Mikha"
		ageMikha   = 23
	)

	// age = 100 --> error

	fmt.Println("Nama ", firstName, "umur ", age)
}
