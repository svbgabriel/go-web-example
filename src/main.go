package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

type Post struct {
    Id    int
    Title string
    Body  string
}

var db, err = sql.Open("mysql", "go:go@/go-web?charset=utf8")

func main() {

    smtm, err := db.Prepare("INSERT INTO posts(title, body) values(?,?)")
    checkErr(err)

    _, err = smtm.Exec("My First Post", "My First Content")
    checkErr(err)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        post := Post{Id: 1, Title: "Unnamed post", Body: "No content"}

        if title := r.FormValue("title"); title != "" {
            post.Title = title
        }

        t := template.Must(template.ParseFiles("templates/index.html"))
        if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    fmt.Println(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
