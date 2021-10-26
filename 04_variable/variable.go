package main

import "fmt"

func main() {
	var name string // di deklarasikan duluan tanpa dikasih tau isinya apa (bikin wadah kosong)
	name = "Duma Doniagara Sambora"

	var age int
	age = 25

	// name = "Eh sekarang diganti"
	// variable bisa ditimpa dengan nilai baru nya

	var friendName = "Mikha" // variable nya langsung diisikan
	var ageFriend = 23

	fmt.Println("Nama saya ", name, "berumur ", age, "temannya adalah ", friendName, "umurnya adalah ", ageFriend)

	// di golang sebenarnya gausah pakai kata kunci var jika langsung di assign
	// contohnya menggunakan titik dua sama dengan (":=")

	country := "Indonesia"
	fmt.Println("Negaranya ", name, "adalah ", country)

	// Multiple declaration
	var (
		firstName = "Duma"
		lastName  = "Sambora"
	)
	fmt.Println(firstName, lastName)
}
