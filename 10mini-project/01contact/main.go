package main

import (
	"01contact/types"
	"fmt"
)

func main() {
	fmt.Println("Contact")
	user := types.Contact{Firstname: "Subhasish", Lastname: "Biswasray", Code: 91, Phone: 9583672357, Email: "biswasray@gmail.com", Address: types.Address{FlatNumber: 1, House: "SriMaa bhavan", Street: "Adhangadhia", ZipCode: 754001, City: "Cuttack", State: "Odisha", Country: "India"}}
	fmt.Println(user)

}
