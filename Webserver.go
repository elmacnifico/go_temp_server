package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "log"
)

type Output struct {
    host string
    DataPoints [60][60]float64
}

type Webserver struct {
    cache *Cache
}

func (self *Webserver) StartServer( cache *Cache ) {
    self.cache = cache
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    host := r.URL.Path[1:]
    output := &Output{host:host}
    for i:=0; i < 60; i++ {
        for j:=0; j < 60; j++ {
            output.DataPoints[i][j] = web.cache.Content[host].Minutes[i].Seconds[j]
        }
    }
    b, err := json.Marshal(output)
    if err != nil {
        log.Println(err)
    }
    fmt.Fprintf(w, "%s", b)
}
