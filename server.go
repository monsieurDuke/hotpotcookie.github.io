package main

import (
    "log"
    "net/http"
    "cookie/source/golang"
)
// main function
func main() {    
    url := "http://localhost:8080"    
    mux := http.NewServeMux()    
    log.Println("[gfarms]: starting server on",url)
    //    
    mux.HandleFunc("/",golang.HomePage)    
    mux.HandleFunc("/dashboard",golang.DboardPage)    
    //
    mux.Handle("/source/css/", http.StripPrefix("/source/css/", http.FileServer(http.Dir("source/css"))))
    mux.Handle("/source/js/", http.StripPrefix("/source/js/", http.FileServer(http.Dir("source/js"))))    
    err := http.ListenAndServe(":8080",mux)
    log.Fatal(err)
}