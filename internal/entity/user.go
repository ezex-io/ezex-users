package entity

type CreateUserRequest struct {
	FirebaseUID string
	Email       string
}

type CreateUserResponse struct {
	UserID string
}

type GetUserByEmailRequest struct {
	Email string
}

type GetUserByEmailResponse struct {
	ID             string
	Email          string
	FirebaseUID    string
	SecurityImage  string
	SecurityPhrase string
}
