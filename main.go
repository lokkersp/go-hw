package main

import (
    "fmt"
    "os"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "shalom")
}
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(fmt.Sprintf(":%s",getEnv("L_PORT","8080")), nil)
}
