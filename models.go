package main

//Principal Object
type Cloth struct {
	ID       int    `json:"id"`
	Image    string `json:"image"`    
	Category string `json:"category"` 
	Style    string `json:"style"`    
}

// Structures for Comunication with Python

type AIRequest struct {
	Image string `json:"image"`
}


type AIResponse struct {
	Status   string `json:"status"`
	Analysis struct {
		Category   string  `json:"category"`
		Style      string  `json:"style"`
		Confidence float64 `json:"confidence"`
	} `json:"analysis"`
}