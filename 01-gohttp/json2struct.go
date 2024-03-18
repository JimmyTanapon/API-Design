package main

import (
	"encoding/json"
	"fmt"
)

type User2 struct {
	ID   int
	Name string
	Age  int
}

func testJsontoStruct() {

	data := []byte(`{
		"id":4,
		"name":"Jittanan",
		"age":"24"
	}`)
	u := &User2{}
	err := json.Unmarshal(data, u)
	fmt.Println("%#v \n", u)
	fmt.Println(err)

}
