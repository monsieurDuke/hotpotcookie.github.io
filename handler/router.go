package handler

import (
    "log"
    "net/http"
    "path"
    "html/template"
)
var url string = "http://localhost:8080"

// page router
func HomePage(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w,r)
        log.Println("[gfarms]: req HTTP/404 "+url+r.URL.Path,r.RemoteAddr)
        return
    }
//    w.Write([]byte("Home Page"))
    tmp,err := template.ParseFiles(path.Join("pages","index.html"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return        
    }
    tmp.Execute(w,nil)
    log.Println("[gfarms]: req HTTP/1.1 "+url+"/",r.RemoteAddr) // bikin log kalo error
}
func DboardPage(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Dashboard Page"))
    log.Println("[gfarms]: req HTTP/1.1 "+url+"/dashboard",r.RemoteAddr)
}