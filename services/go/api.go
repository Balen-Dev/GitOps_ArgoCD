package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
    http.HandleFunc("/", generalGreetingHandler)
    http.HandleFunc("/greet/", customGreetingHandler)

    fmt.Println("Server is running on http://localhost:8080")
    // log.Fatal(http.ListenAndServe(":8080", nil))
    srv := &http.Server{
        Addr:         ":8080",
        Handler:      nil,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }
    
    log.Fatal(srv.ListenAndServe())    
}

func generalGreetingHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet || r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintln(w, "Hello! Welcome to the GO API. (Deployed by ArgoCD)")
}

func customGreetingHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Extract name from URL
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 || parts[2] == "" {
        http.Error(w, "Name not provided", http.StatusBadRequest)
        return
    }

    name := parts[2]
    greeting := fmt.Sprintf("Hello, %s! (GO)", name)
    fmt.Fprintln(w, greeting)
}
