package events

type UserEvent struct {
	EventType string `json:"event_type"` // "created", "updated", "deleted"
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
}
