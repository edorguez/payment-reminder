package handlers

import "time"

type UserResponse struct {
	ID              int64     `json:"id"`
	FirebaseUID     string    `json:"firebaseUid"`
	UserPlanID      int64     `json:"userPlanId"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	LastPaymentDate time.Time `json:"lastPaymentDate"`
}

type UpdateUserRequest struct {
	UserPlanID int64  `json:"userPlanId" binding:"required"`
	Email      string `json:"email" binding:"required"`
}
