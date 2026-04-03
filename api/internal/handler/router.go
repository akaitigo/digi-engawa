package handler

import (
	"net/http"

	"github.com/akaitigo/digi-engawa/api/internal/repository"
	"github.com/akaitigo/digi-engawa/api/internal/service"
	"github.com/akaitigo/digi-engawa/api/internal/ws"
)

func NewRouter(dataDir string) (*http.ServeMux, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)

	hub := ws.NewHub()

	materialRepo, err := repository.NewMaterialRepository(dataDir)
	if err != nil {
		return nil, err
	}
	materialSvc := service.NewMaterialService(materialRepo)
	materialHandler := NewMaterialHandler(materialSvc)
	materialHandler.Register(mux)

	helpRepo, err := repository.NewHelpRequestRepository(dataDir)
	if err != nil {
		return nil, err
	}
	helpSvc := service.NewHelpRequestService(helpRepo, hub)
	helpHandler := NewHelpRequestHandler(helpSvc)
	helpHandler.Register(mux)

	classroomRepo, err := repository.NewClassroomRepository(dataDir)
	if err != nil {
		return nil, err
	}
	classroomSvc := service.NewClassroomService(classroomRepo)
	classroomHandler := NewClassroomHandler(classroomSvc)
	classroomHandler.Register(mux)

	progressRepo, err := repository.NewProgressRepository(dataDir)
	if err != nil {
		return nil, err
	}
	progressHandler := NewProgressHandler(progressRepo, classroomRepo, hub)
	progressHandler.Register(mux)

	return mux, nil
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
