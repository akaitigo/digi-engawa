package model

import "time"

type LearnerProgress struct {
	ID            string    `json:"id"`
	ParticipantID string    `json:"participant_id"`
	MaterialID    string    `json:"material_id"`
	CurrentStep   int       `json:"current_step"`
	Completed     bool      `json:"completed"`
	UpdatedAt     time.Time `json:"updated_at"`
}
