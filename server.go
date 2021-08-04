package main

import (
    "log"
    "net/http"
    "cookie/handler"
)
// main function
func main() {    
    url := "http://localhost:8080"    
    mux := http.NewServeMux()    
    log.Println("[gfarms]: starting server on",url)
    //    
    mux.HandleFunc("/",handler.HomePage)    
    mux.HandleFunc("/dashboard",handler.DboardPage)    
    err := http.ListenAndServe(":8080",mux)
    log.Fatal(err)
}