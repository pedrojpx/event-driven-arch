package web

import (
	"encoding/json"
	"net/http"

	createaccount "github.com/pedrojpx/ms-wallet/internal/usecase/create_account"
)

type WebAccountHandler struct {
	createUseCase createaccount.CreateAccountUseCase
}

func NewWebAccountHandler(uc createaccount.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		createUseCase: uc,
	}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto createaccount.CreateAccountInputDTO
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
