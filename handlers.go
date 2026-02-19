package main

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "strconv"
)

var closet []Cloth  // In-memory database

var nextID = 1      // Counter

// CORS Permissions
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// GET: Retrieve all clothes
func GetClothes(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(closet) // Send JSON response
}

// POST: Add a new cloth
func CreateCloth(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    if r.Method == "OPTIONS" { return } // Handle preflight request

    var newCloth Cloth
    json.NewDecoder(r.Body).Decode(&newCloth) // Read request body

    // Call for python service
    pyReq := AIRequest{Image: newCloth.Image}
    jsonData, _ := json.Marshal(pyReq) // Prepare JSON payload
    resp, err := http.Post("http://127.0.0.1:5000/analizar", "application/json", bytes.NewBuffer(jsonData))
    
    if err == nil { // If Python service answers successfully
        defer resp.Body.Close()
        body, _ := io.ReadAll(resp.Body)
        var aiResp AIResponse
        json.Unmarshal(body, &aiResp) // Parse AI response
        
        // Assign AI results to the new cloth
        newCloth.Category = aiResp.Analysis.Category
        newCloth.Style = aiResp.Analysis.Style
    }

    // Assign ID and save to slice
    newCloth.ID = nextID
    nextID++
    closet = append(closet, newCloth)
    
    // Return created cloth
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newCloth)
}

// UPDATE: Edit existing cloth
func UpdateCloth(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    if r.Method == "OPTIONS" { return } // Handle preflight request

    // Get ID from URL query
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    var updatedData Cloth
    json.NewDecoder(r.Body).Decode(&updatedData) // Read new data

    // Find cloth by ID
    for i, cloth := range closet {
        if cloth.ID == id {
            // Update only if new data is provided
            if updatedData.Category != "" { closet[i].Category = updatedData.Category }
            if updatedData.Style != "" { closet[i].Style = updatedData.Style }
            
            // Return updated cloth
            json.NewEncoder(w).Encode(closet[i])
            return
        }
    }
}