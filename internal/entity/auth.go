package entity

type ProcessFirebaseLoginRequest struct {
	Email       string
	FirebaseUID string
}

type ProcessFirebaseLoginResponse struct {
	UserID string
}
