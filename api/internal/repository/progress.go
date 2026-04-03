package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/akaitigo/digi-engawa/api/internal/model"
)

type ProgressRepository struct {
	mu      sync.RWMutex
	dataDir string
	progress map[string]model.LearnerProgress // key: participantID:materialID
}

func NewProgressRepository(dataDir string) (*ProgressRepository, error) {
	r := &ProgressRepository{
		dataDir:  dataDir,
		progress: make(map[string]model.LearnerProgress),
	}
	if err := r.loadFromDisk(); err != nil {
		return nil, fmt.Errorf("load progress: %w", err)
	}
	return r, nil
}

func (r *ProgressRepository) Upsert(p model.LearnerProgress) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := p.ParticipantID + ":" + p.MaterialID
	r.progress[key] = p
	return r.saveToDisk()
}

func (r *ProgressRepository) Get(participantID, materialID string) (model.LearnerProgress, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.progress[participantID+":"+materialID]
	return p, ok
}

func (r *ProgressRepository) GetByClassroom(classroomParticipantIDs []string) []model.LearnerProgress {
	r.mu.RLock()
	defer r.mu.RUnlock()

	idSet := make(map[string]bool)
	for _, id := range classroomParticipantIDs {
		idSet[id] = true
	}

	var result []model.LearnerProgress
	for _, p := range r.progress {
		if idSet[p.ParticipantID] {
			result = append(result, p)
		}
	}
	return result
}

func (r *ProgressRepository) dataFilePath() string {
	return filepath.Join(r.dataDir, "progress.json")
}

func (r *ProgressRepository) loadFromDisk() error {
	data, err := os.ReadFile(r.dataFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var items []model.LearnerProgress
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	for _, p := range items {
		r.progress[p.ParticipantID+":"+p.MaterialID] = p
	}
	return nil
}

func (r *ProgressRepository) saveToDisk() error {
	if err := os.MkdirAll(r.dataDir, 0o755); err != nil {
		return err
	}

	items := make([]model.LearnerProgress, 0, len(r.progress))
	for _, p := range r.progress {
		items = append(items, p)
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.dataFilePath(), data, 0o644)
}
