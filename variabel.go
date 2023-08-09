package main

import "fmt"

func main(){
	// variabel
	var lastName string = "Geofani"
	var firstName string = "Hana"
	
	//konstanta
	// firstName = "anah"
	// const lastName = "Geofani"

	//datatypes
	// var bilanganBulat uint8 = 20
	// var bilanganDesimal = 2.0
	// var varBool = true

	//variabel pointer
	var numberA int = 4
	var numberB *int = &numberA
	*numberB = 8 

	fmt.Println("Hai", firstName, lastName, numberA, "!\n")
}