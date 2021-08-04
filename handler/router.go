package handler

import (
    "log"
    "net/http"
)
var url string = "http://localhost:8080"

// page router
func HomePage(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w,r)
        log.Println("[gfarms]: req HTTP/404 "+url+r.URL.Path,r.RemoteAddr)
        return
    }
    w.Write([]byte("Home Page"))
    log.Println("[gfarms]: req HTTP/1.1 "+url+"/",r.RemoteAddr)
}
func DboardPage(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Dashboard Page"))
    log.Println("[gfarms]: req HTTP/1.1 "+url+"/dashboard",r.RemoteAddr)
}