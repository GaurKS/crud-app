package dtos

// user - DTOs
type CreateTodo struct {
	Title          string        `json:"title,omitempty"`
	TodoStatus     string        `json:"todoStatus,omitempty"`
	Description    string        `json:"description,omitempty"`
	CreatedBy      string        `json:"createdBy,omitempty"`
}