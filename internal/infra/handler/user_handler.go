package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PGabriel20/expenses-go/internal/entity"
	"github.com/PGabriel20/expenses-go/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	UserRepository entity.UserRepository
}

func NewUserHandler(userRepository entity.UserRepository, r *chi.Mux) {
	handler := &UserHandler{
		UserRepository: userRepository,
	}

	//User route group
	r.Route("/v1/user", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Get("/{id}", handler.Get)
	})
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var dto usecase.UserInputDto

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registerUser := usecase.NewRegisterUserUseCase(h.UserRepository)
	user, err := registerUser.ExecuteRegister(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

	return
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")

	getUser := usecase.NewGetUserUseCase(h.UserRepository)
	user, err := getUser.ExecuteGet(uid)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

	return
}