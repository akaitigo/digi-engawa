package handler

import (
	"net/http"

	"github.com/akaitigo/digi-engawa/api/internal/repository"
	"github.com/akaitigo/digi-engawa/api/internal/service"
)

func NewRouter(dataDir string) (*http.ServeMux, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)

	materialRepo, err := repository.NewMaterialRepository(dataDir)
	if err != nil {
		return nil, err
	}
	materialSvc := service.NewMaterialService(materialRepo)
	materialHandler := NewMaterialHandler(materialSvc)
	materialHandler.Register(mux)

	return mux, nil
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
