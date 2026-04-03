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

type ClassroomRepository struct {
	mu         sync.RWMutex
	dataDir    string
	classrooms map[string]model.Classroom
	participants map[string][]model.Participant
}

func NewClassroomRepository(dataDir string) (*ClassroomRepository, error) {
	r := &ClassroomRepository{
		dataDir:      dataDir,
		classrooms:   make(map[string]model.Classroom),
		participants: make(map[string][]model.Participant),
	}
	if err := r.loadFromDisk(); err != nil {
		return nil, fmt.Errorf("load classrooms: %w", err)
	}
	return r, nil
}

func (r *ClassroomRepository) SaveClassroom(c model.Classroom) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.classrooms[c.ID] = c
	return r.saveToDisk()
}

func (r *ClassroomRepository) GetClassroomByID(id string) (model.Classroom, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	c, ok := r.classrooms[id]
	return c, ok
}

func (r *ClassroomRepository) GetClassroomByCode(code string) (model.Classroom, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, c := range r.classrooms {
		if c.ClassroomCode == code {
			return c, true
		}
	}
	return model.Classroom{}, false
}

func (r *ClassroomRepository) GetAllClassrooms() []model.Classroom {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.Classroom, 0, len(r.classrooms))
	for _, c := range r.classrooms {
		result = append(result, c)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ScheduledAt.After(result[j].ScheduledAt)
	})
	return result
}

func (r *ClassroomRepository) AddParticipant(p model.Participant) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.participants[p.ClassroomID] = append(r.participants[p.ClassroomID], p)
	return r.saveToDisk()
}

func (r *ClassroomRepository) GetParticipants(classroomID string) []model.Participant {
	r.mu.RLock()
	defer r.mu.RUnlock()

	src := r.participants[classroomID]
	if len(src) == 0 {
		return []model.Participant{}
	}
	result := make([]model.Participant, len(src))
	copy(result, src)
	return result
}

type classroomData struct {
	Classrooms   []model.Classroom              `json:"classrooms"`
	Participants map[string][]model.Participant `json:"participants"`
}

func (r *ClassroomRepository) dataFilePath() string {
	return filepath.Join(r.dataDir, "classrooms.json")
}

func (r *ClassroomRepository) loadFromDisk() error {
	data, err := os.ReadFile(r.dataFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var d classroomData
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	for _, c := range d.Classrooms {
		r.classrooms[c.ID] = c
	}
	if d.Participants != nil {
		r.participants = d.Participants
	}
	return nil
}

func (r *ClassroomRepository) saveToDisk() error {
	if err := os.MkdirAll(r.dataDir, 0o755); err != nil {
		return err
	}

	classrooms := make([]model.Classroom, 0, len(r.classrooms))
	for _, c := range r.classrooms {
		classrooms = append(classrooms, c)
	}

	d := classroomData{
		Classrooms:   classrooms,
		Participants: r.participants,
	}

	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.dataFilePath(), data, 0o644)
}
