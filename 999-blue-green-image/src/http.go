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

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I am %s\n", hostname)

        // Print out the hostname with background-color generated from it
        fmt.Fprintf(w, "<body style='background-color:green\n")
        fmt.Fprintf(w, ";'>")
        fmt.Fprintf(w, "<p style='color: #000000; background-color: #ffffff; font-weight: bold'>I am: %s\n", hostname)
        fmt.Fprintf(w, "</p></body>")
    })
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
