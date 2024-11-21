package web

import (
	"encoding/json"
	"net/http"

	createtransaction "github.com/pedrojpx/ms-wallet/internal/usecase/create_transaction"
)

type WebTransactionHandler struct {
	createUseCase createtransaction.CreateTransactionUsecase
}

func NewWebTransactionHandler(uc createtransaction.CreateTransactionUsecase) *WebTransactionHandler {
	return &WebTransactionHandler{
		createUseCase: uc,
	}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto createtransaction.CreateTrasactionInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output, err := h.createUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
