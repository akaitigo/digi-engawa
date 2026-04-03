package model

import "time"

type HelpRequest struct {
	ID             string     `json:"id"`
	ClassroomID    string     `json:"classroom_id"`
	ParticipantID  string     `json:"participant_id"`
	MaterialStepID string     `json:"material_step_id"`
	Status         string     `json:"status"`
	CreatedAt      time.Time  `json:"created_at"`
	ResolvedAt     *time.Time `json:"resolved_at,omitempty"`
}
