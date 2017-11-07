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

        out_config, err_config := exec.Command("cat", "/app/config/conf.txt").Output()
        config := ""
        if err != nil {
            config = err_config.Error();
        } else {
            config = string(out_config);
        }

        out_count, err_count := exec.Command("sh", "/app/counter.sh").Output()
        count := ""
        if err != nil {
            count = err_count.Error();
        } else {
            count = string(out_count);
        }

        fmt.Fprintf(os.Stdout,  "I'm %s - backend response : %s - config : %s - count : %s\n", hostname, body, out_config, out_count)
        fmt.Fprintf(w,          "I'm %s - backend response : %s\n- config : %s\n- count : %s\n", hostname, body, out_config, out_count)
        fmt.Fprintf(os.Stdout,  "config : %s - count : %s\n", config, count)
    })



    log.Fatal(http.ListenAndServe(":" + port, nil))
}

