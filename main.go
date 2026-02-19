package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Register route handlers
    http.HandleFunc("/clothes", GetClothes)   // GET: List inventory
    http.HandleFunc("/add", CreateCloth)      // POST: Add new item
    http.HandleFunc("/update", UpdateCloth)   // UPDATE: Edit existing item

    // Start server message
    fmt.Println("Go corriendo en http://localhost:8080")
    
    // Listen on port 8080 and log errors if it crashes
    log.Fatal(http.ListenAndServe(":8080", nil))
}