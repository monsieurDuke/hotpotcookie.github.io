package main

import (
    "fmt"
    //"log"
    "net/http"
)
// MAIN METHOD
func main() {
    fmt.Println("[grow-farms]: serving website in :8080")	
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/contact", contactHandler)    
    http.ListenAndServe(":8080", nil)
}

// PAGE HANDLER
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    fmt.Fprint(w, "SERVING HELLO")
    fmt.Println("[grow-farms]: responding",r.URL.Path,"(HTTP/1.1)")    
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/contact" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    fmt.Fprint(w, "SERVING HELLO")
    fmt.Println("[grow-farms]: responding",r.URL.Path,"(HTTP/1.1)")    
}

// ERROR HANDLER
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "SERVING ERROR 404")
	    fmt.Println("[grow-farms]: responding",r.URL.Path,"(HTTP/404)")
    }
}