//easycode.go serves a webpage useful for showing a demo of Cloud Foundry
package main

import (
	"net/http"
	"html/template"
	"log"
	"encoding/json"
)


type Result struct {
	Name string
	Value int
} 

type Message struct {
	Text string
}

func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}


func showTemplate(w http.ResponseWriter, r *http.Request) {
	var results=[]Result {
				{Name:"Choice 1", Value:10},
				{Name:"Choice 2", Value:30},
			}
	t,err:=template.ParseFiles("templates/index.tmpl")
	check("Parse template",err)

	err=t.Execute(w,results)
	if (err != nil) {
		log.Println("Error: ", err)
	}
}

func displayAbout (w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome to the EasyAPI, build v0.0.001.001"}
    b, err := json.Marshal(m)
 
    if err != nil {
        panic(err)
    }
 
     w.Write(b)
}


func main() {
	http.HandleFunc("/",showTemplate)
	http.HandleFunc("/about",displayAbout)
	http.Handle("/images/",http.FileServer(http.Dir("")))
	http.Handle("/css/",http.FileServer(http.Dir("")))
	http.Handle("/fonts/",http.FileServer(http.Dir("")))
	http.Handle("/js/",http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(":8080",nil))
}