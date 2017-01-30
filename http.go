package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "os/exec"
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


    fmt.Fprintf(os.Stdout, "Listening on :%s, call %s\n", port, backEndURL)


    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

               
        out, err := exec.Command("curl", backEndURL).Output()
        body := ""
        if err != nil {
            body = err.Error();
        } else {
            body = string(out);
        }


        fmt.Fprintf(os.Stdout, "I'm %s - backend response : %s \n", hostname, body)
 	    fmt.Fprintf(w, "I'm %s - backend response : %s \n", hostname, body)
    })



    log.Fatal(http.ListenAndServe(":" + port, nil))
}

