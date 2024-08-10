package models

// Task struct represents the structure of a task
type Task struct {
	ID          string `json:"id"bson:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}
