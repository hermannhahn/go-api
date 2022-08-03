package models

// Response struct
type Response struct {
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

// ResponseList struct
type ResponseList struct {
	Message string   `json:"message"`
	Data    Products `json:"data"`
}
