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

type MaterialRepository struct {
	mu       sync.RWMutex
	dataDir  string
	materials map[string]model.Material
}

func NewMaterialRepository(dataDir string) (*MaterialRepository, error) {
	r := &MaterialRepository{
		dataDir:  dataDir,
		materials: make(map[string]model.Material),
	}
	if err := r.loadFromDisk(); err != nil {
		return nil, fmt.Errorf("load materials: %w", err)
	}
	return r, nil
}

func (r *MaterialRepository) GetAll() []model.Material {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.Material, 0, len(r.materials))
	for _, m := range r.materials {
		result = append(result, m)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.Before(result[j].CreatedAt)
	})
	return result
}

func (r *MaterialRepository) GetByID(id string) (model.Material, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	m, ok := r.materials[id]
	return m, ok
}

func (r *MaterialRepository) Save(m model.Material) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.materials[m.ID] = m
	return r.saveToDisk()
}

func (r *MaterialRepository) dataFilePath() string {
	return filepath.Join(r.dataDir, "materials.json")
}

func (r *MaterialRepository) loadFromDisk() error {
	data, err := os.ReadFile(r.dataFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var materials []model.Material
	if err := json.Unmarshal(data, &materials); err != nil {
		return err
	}

	for _, m := range materials {
		r.materials[m.ID] = m
	}
	return nil
}

func (r *MaterialRepository) saveToDisk() error {
	if err := os.MkdirAll(r.dataDir, 0o755); err != nil {
		return err
	}

	materials := make([]model.Material, 0, len(r.materials))
	for _, m := range r.materials {
		materials = append(materials, m)
	}

	data, err := json.MarshalIndent(materials, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.dataFilePath(), data, 0o644)
}
