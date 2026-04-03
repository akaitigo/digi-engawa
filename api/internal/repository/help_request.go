package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/akaitigo/digi-engawa/api/internal/model"
)

type HelpRequestRepository struct {
	mu      sync.RWMutex
	dataDir string
	requests map[string]model.HelpRequest
}

func NewHelpRequestRepository(dataDir string) (*HelpRequestRepository, error) {
	r := &HelpRequestRepository{
		dataDir:  dataDir,
		requests: make(map[string]model.HelpRequest),
	}
	if err := r.loadFromDisk(); err != nil {
		return nil, fmt.Errorf("load help requests: %w", err)
	}
	return r, nil
}

func (r *HelpRequestRepository) Save(hr model.HelpRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.requests[hr.ID] = hr
	return r.saveToDisk()
}

func (r *HelpRequestRepository) GetByID(id string) (model.HelpRequest, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	hr, ok := r.requests[id]
	return hr, ok
}

func (r *HelpRequestRepository) GetByClassroom(classroomID string) []model.HelpRequest {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []model.HelpRequest
	for _, hr := range r.requests {
		if hr.ClassroomID == classroomID {
			result = append(result, hr)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.Before(result[j].CreatedAt)
	})
	return result
}

func (r *HelpRequestRepository) dataFilePath() string {
	return filepath.Join(r.dataDir, "help_requests.json")
}

func (r *HelpRequestRepository) loadFromDisk() error {
	data, err := os.ReadFile(r.dataFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var requests []model.HelpRequest
	if err := json.Unmarshal(data, &requests); err != nil {
		return err
	}

	for _, hr := range requests {
		r.requests[hr.ID] = hr
	}
	return nil
}

func (r *HelpRequestRepository) saveToDisk() error {
	requests := make([]model.HelpRequest, 0, len(r.requests))
	for _, hr := range r.requests {
		requests = append(requests, hr)
	}

	data, err := json.MarshalIndent(requests, "", "  ")
	if err != nil {
		return err
	}

	return atomicWriteFile(r.dataFilePath(), data, 0o600)
}
