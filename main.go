package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/contact.html")
}

func main() {
    fmt.Println("Starting server on port 9090...")

    http.HandleFunc("/home", homePage)
    http.HandleFunc("/courses", coursePage)
    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/contact", contactPage)

    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("Server error:", err)
    }
}
