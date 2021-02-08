package models

//User structure
type Success struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

var (
	InsertSuccess = Success{Success: 1, Message: "insert success"}
)
