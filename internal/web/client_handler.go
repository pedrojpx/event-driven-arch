package web

import (
	"encoding/json"
	"net/http"

	createclient "github.com/pedrojpx/ms-wallet/internal/usecase/create_client"
)

type WebClientHandler struct {
	createUseCase createclient.CreateClientUseCase
}

func NewWebClientHandler(uc createclient.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		createUseCase: uc,
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto createclient.CreateClientInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	output, err := h.createUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
