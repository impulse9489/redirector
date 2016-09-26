package main

import (
    "log"
    "net/http"
    "os"
    "strings"
    "fmt"
)



func redirect(w http.ResponseWriter, r *http.Request) {
    if len(os.Args) != 3 {
        http.Error(w, "Usage: redirect listener_port redirect_url", 400)
        return
    }
    redirect_site := os.Args[2]
    redirect_no_schema := strings.Replace(redirect_site,"http://","",1)
    redirect_no_schema = strings.Replace(redirect_no_schema,"https://","",1)
    if redirect_no_schema == r.Host {
        http.Error(w, "Sorry not allowed as this will cause an infinite loop.", 400)
    } else {
        http.Redirect(w, r, strings.TrimRight(redirect_site,"/") + r.URL.Path, 301)
    }
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: redirect listener_port redirect_url\nEx: redirect 80 http://www.tufts.edu")
        return
    }
    listener_port := os.Args[1]
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":" + listener_port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    } 
}