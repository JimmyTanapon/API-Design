package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ข้อควรรู้ตัวเเปรในstruct ถ้าเป็นตัวเล็กจะเป็น private

var user = []User{
	{ID: 1, Name: "Anuchito", Age: 22},
	{ID: 2, Name: "Jimmy", Age: 28},
	{ID: 3, Name: "weerakul23May", Age: 32},
}

func handle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		b, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(b)
		return
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var u User

		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		user = append(user, u)
		fmt.Fprintf(w, "hello %s created users", "POST")
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func main() {
	http.HandleFunc("/users", handle)

	log.Println("Server start at port:2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
	log.Println("bye....")
}
