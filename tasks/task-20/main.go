package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"demo/handlers"
)

var (
	PORT string
)

func main() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "9091", "--port=9091 or -port=9091")
	}
	flag.Parse()

	fmt.Println("Server started and listening on port ", PORT)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	http.HandleFunc("/greet", SayHi)
	http.HandleFunc("/greetby", GreetBy)
	ehandler:=&handlers.EmployeeHandler{}
	http.HandleFunc("/employee/add",ehandler.AddEmployee)
	http.HandleFunc("/employee/delete", ehandler.DeleteEmployee)


	

	err:=http.ListenAndServe("0.0.0.0:"+PORT,nil)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(err,"Hello World")


}

func SayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Victoria Secrets & Co")
}

func GreetBy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Hit here?")
		g := Greet{}
		err := json.NewDecoder(r.Body).Decode(&g) //takes data from response body and saves it to the struct g
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		} else {
			fmt.Fprintln(w, "Data has been read", g)
		}
	} else {
		w.Write([]byte("Unimplemented Method"))
		w.WriteHeader(http.StatusNotImplemented)
	}

}

type Greet struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}
