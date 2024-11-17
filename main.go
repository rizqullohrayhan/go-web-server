package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var name = "John"
type prodSpec struct {
	Size	string
	Weight	float32
	Descr	string
}

type produc struct {
	ID		int32
	Name	string
	Cost	float64
	Specs	prodSpec
}

type User struct  {
	Name		string
	Language	string
	Member		bool
}

var tpl *template.Template
var prodi produc
var user User

func main() {
	prodi = produc{
		ID: 15,
		Name: "Samsung Galaxy S25 Ultra",
		Cost: 500,
		Specs: prodSpec{
			Size: "150 x 70 x 7 mm",
			Weight: 65,
			Descr: "Over priced shiny thing designed to shatter on impact",
		},
	}

	user = User{"Rizqulloh Rayhan", "Indonesia", false}

	tpl, _ = tpl.ParseGlob("view/*.html")

	http.HandleFunc("/", indexHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)
	http.HandleFunc("/contact", contactHandleFunc)
	http.HandleFunc("/login", loginHandleFunc)
	http.HandleFunc("/productinfo", productInfoHandler)
	http.HandleFunc("/membership", membershipHandler)

	fmt.Println("listening to 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandleFunc(w http.ResponseWriter, r *http.Request)  {
	tpl.ExecuteTemplate(w, "index.html", name)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request)  {
	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandleFunc(w http.ResponseWriter, r *http.Request)  {
	tpl.ExecuteTemplate(w, "contact.html", nil)
}

func loginHandleFunc(w http.ResponseWriter, r *http.Request)  {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func productInfoHandler(w http.ResponseWriter, r *http.Request)  {
	tpl.ExecuteTemplate(w, "productInfo.html", prodi)
}

func membershipHandler(w http.ResponseWriter, r *http.Request)  {
	tpl.ExecuteTemplate(w, "membership.html", user)
}