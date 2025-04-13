package handler

import (
	"cv-api/internal/usecase"
	"encoding/json"
	"html/template"
	"net/http"
)

type CVHandler struct {
	cvUsecase usecase.CVUsecase
}

func NewCVHandler(cvUsecase usecase.CVUsecase) *CVHandler {
	return &CVHandler{
		cvUsecase: cvUsecase,
	}
}

func (h *CVHandler) ServeHTML(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ajout des en-têtes CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	tmpl.Execute(w, nil)
}

func (h *CVHandler) GetCV(w http.ResponseWriter, r *http.Request) {
	// Ajout des en-têtes CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	cv, err := h.cvUsecase.GetCV()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cv)
}
