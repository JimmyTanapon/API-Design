package main

import (
	"encoding/json"
	"fmt"
)

// เเปลง stuct ไปเป็น json
func testStuctToJson() {
	u := User{
		ID: 1, Name: "Anuchito", Age: 22,
	}
	user, err := json.Marshal(u)
	fmt.Printf("byte : %s \n ", user)
	fmt.Println(err)

}
