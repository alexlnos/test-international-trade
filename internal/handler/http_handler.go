package handler

import (
	"encoding/json"
	"net/http"
	"test-international-trade/internal/usecase"
)

// HTTPHandler handles HTTP requests for the encryption service
type HTTPHandler struct {
	useCase *usecase.EncryptUseCase
}

func NewHTTPHandler(useCase *usecase.EncryptUseCase) *HTTPHandler {
	return &HTTPHandler{useCase: useCase}
}

// Encrypt godoc
// @Summary Encrypt a string
// @Description Encrypts a string using MD5 and SHA256 algorithms
// @Tags encryption
// @Accept json
// @Produce json
// @Param input body domain.EncryptRequest true "Input string to encrypt"
// @Success 200 {object} domain.EncryptedData
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /encrypt [post]
func (h *HTTPHandler) Encrypt(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Input string `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := h.useCase.Encrypt(req.Input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
