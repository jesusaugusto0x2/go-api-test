package input

type CreateUserInput struct {
	Name  string
	Email string
}

type UpdateUserInput struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}
