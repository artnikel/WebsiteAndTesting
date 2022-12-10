package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Mail     string
	Password string
}

var user Users

var rootMail string
var rootPassword string
var allrootUsers string

func signin_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/signin.html")
	tmpl.ExecuteTemplate(w, "signin", nil)
}

func database_insert(request string) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query(request)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func check_login(email string, password string, w http.ResponseWriter, r *http.Request) bool {
	var check bool
	if email == "artyom.leskco@gmail.com" && password == "123" || email == "root@root" && password == "root" {
		check = true
		database_insert(fmt.Sprintf("INSERT INTO `users`(`Mail`, `Password`) VALUES ('%s','%s')", email, password))
		http.Redirect(w, r, "/home/", http.StatusSeeOther)
	} else {
		check = false
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	return check
}

func authorize(w http.ResponseWriter, r *http.Request) {
	check_login(r.FormValue("email"), r.FormValue("password"), w, r)
}

func database_select(request string) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query(request)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err = results.Scan(&user.Mail, &user.Password)
		if err != nil {
			panic(err.Error())
		}
	}
	defer results.Close()
}

func home_page(w http.ResponseWriter, r *http.Request) {

	database_select("SELECT `Mail`,`Password` FROM `users` ORDER BY `ID` DESC LIMIT 1")
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(w, user.Mail)
}

func database_alldata() {
	allrootUsers = ""
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT `Mail`,`Password` FROM `users`")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err = results.Scan(&rootMail, &rootPassword)
		if err != nil {
			panic(err.Error())
		}
		allrootUsers = allrootUsers + "|" + rootMail + "(" + rootPassword + ")"
	}
	defer results.Close()
}

func root_page(w http.ResponseWriter, r *http.Request) {
	database_alldata()
	tmpl, _ := template.ParseFiles("templates/root.html")
	tmpl.ExecuteTemplate(w, "root", allrootUsers)
}

func transporter(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/transporter.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "transporter", nil)
}

func transit(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/transit.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "transit", nil)
}

func fm(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/fm.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "fm", nil)
}

func crafter(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/crafter.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "crafter", nil)
}

func sprinter(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/sprinter.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "sprinter", nil)
}

func cf85(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/cf85.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "cf85", nil)
}

func master(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/master.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "master", nil)
}

func gaz(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/2705.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "2705", nil)
}

func tga(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/trucks/tga.html", "templates/general/footer.html", "templates/general/header.html")
	tmpl.ExecuteTemplate(w, "tga", nil)
}

func handleRequest() {
	http.HandleFunc("/", signin_page)
	http.HandleFunc("/authorize/", authorize)
	http.HandleFunc("/home/", home_page)
	http.HandleFunc("/root/", root_page)
	http.HandleFunc("/transporter/", transporter)
	http.HandleFunc("/transit/", transit)
	http.HandleFunc("/fm/", fm)
	http.HandleFunc("/crafter/", crafter)
	http.HandleFunc("/sprinter/", sprinter)
	http.HandleFunc("/cf85/", cf85)
	http.HandleFunc("/master/", master)
	http.HandleFunc("/2705/", gaz)
	http.HandleFunc("/tga/", tga)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
