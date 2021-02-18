package models

//User structure
type Success struct {
	Success bool    `json:"success"`
	Message string `json:"message"`
}

var (
	InsertSuccess = Success{Success: true, Message: "insert success"}
)
