package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func userHandler(w http.ResponseWriter, r *http.Request) {

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
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is OK !"))
		return
	}
}

// func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		next.ServeHTTP(w, r)
// 		log.Printf("Server http middleware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
// 	}

// }

type Logger struct {
	Hanler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	l.Hanler.ServeHTTP(w, r)
	log.Printf("Server http middleware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", (userHandler))
	mux.HandleFunc("/health", (healthHandler))

	logMux := Logger{
		Hanler: mux,
	}

	svr := http.Server{
		Addr:    ":2565",
		Handler: logMux,
	}

	log.Println("Server start at port:2565")
	log.Fatal(svr.ListenAndServe())
	log.Println("bye....")
}
