package model

import "time"

type Material struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Steps       []Step    `json:"steps,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Step struct {
	ID           string    `json:"id"`
	MaterialID   string    `json:"material_id"`
	StepOrder    int       `json:"step_order"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	FuriganaBody string    `json:"furigana_body"`
	AudioText    string    `json:"audio_text"`
	CreatedAt    time.Time `json:"created_at"`
}
