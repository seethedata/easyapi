//easycode.go serves a webpage useful for showing a demo of Cloud Foundry
package main

import (
	"net/http"
	"html/template"
	"log"
)


type Result struct {
	Name string
	Value int
} 



func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}


func responseHandler(w http.ResponseWriter, r *http.Request) {
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


func main() {
	http.HandleFunc("/",responseHandler)
	http.Handle("/images/",http.FileServer(http.Dir("")))
	http.Handle("/css/",http.FileServer(http.Dir("")))
	http.Handle("/fonts/",http.FileServer(http.Dir("")))
	http.Handle("/js/",http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(":8080",nil))
}