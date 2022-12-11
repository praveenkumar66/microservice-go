package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hi JERRY")
		d,_ := ioutil.ReadAll(r.Body)

		fmt.Fprintf(w,"Hello %s\n",d)

		//fmt.Printf("Data %s\n",d)
	})
	http.HandleFunc("/goodnight",func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodnight JERRY")
	})

	http.ListenAndServe(":9090",nil)
}