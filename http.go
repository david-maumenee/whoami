package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    backEndURL = os.Getenv("BACK_END_URL")
    if backEndURL == "" {
        backEndURL = "http://backend"
    }


    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)


    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        resp, err := http.Get(backEndURL)

        fmt.Fprintf(os.Stdout, "I'm %s - backend response : %s \n", hostname, resp)
 	    fmt.Fprintf(w, "I'm %s - backend response : %s \n", hostname, resp)
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}

