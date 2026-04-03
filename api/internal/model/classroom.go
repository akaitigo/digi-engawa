package model

import "time"

type Classroom struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Location      string    `json:"location"`
	Capacity      int       `json:"capacity"`
	ScheduledAt   time.Time `json:"scheduled_at"`
	ClassroomCode string    `json:"classroom_code"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Participant struct {
	ID          string    `json:"id"`
	ClassroomID string    `json:"classroom_id"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}
