package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    backEndURL := os.Getenv("BACK_END_URL")
    if backEndURL == "" {
        backEndURL = "http://backend"
    }


    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)


    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        resp, _ := http.Get(backEndURL)
        
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        fmt.Fprintf(os.Stdout, "I'm %s - backend response : %s | %s \n", body)
 	    fmt.Fprintf(w, "I'm %s - backend response : %s | %s \n", hostname, body)
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}

