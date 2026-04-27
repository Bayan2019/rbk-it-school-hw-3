package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Bayan2019/rbk-it-school-hw-3/internal/domain"
	"github.com/go-chi/chi/v5"
)

type userService interface {
	Create(ctx context.Context, input domain.CreateUserInput) (domain.User, error)
	List(ctx context.Context, filter domain.ListUsersFilter) ([]domain.User, error)
	GetByID(ctx context.Context, id int64, includeDeleted bool) (domain.User, error)
	Update(ctx context.Context, id int64, input domain.UpdateUserInput) (domain.User, error)
	Delete(ctx context.Context, id int64) error
}

type UserHandler struct {
	service userService
}

type usersResponse struct {
	Data []domain.User `json:"data"`
}

type userResponse struct {
	Data domain.User `json:"data"`
}

func NewUserHandler(service userService) *UserHandler {
	return &UserHandler{service: service}
}

////// methods
////// methods
////// methods

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid json body"})
		return
	}

	user, err := h.service.Create(r.Context(), input)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Location", "/api/v1/users/"+strconv.FormatInt(user.ID, 10))
	writeJSON(w, http.StatusCreated, userResponse{Data: user})
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	filter := domain.ListUsersFilter{
		Limit:          parseIntQuery(r, "limit", 20),
		Offset:         parseIntQuery(r, "offset", 0),
		Query:          r.URL.Query().Get("q"),
		IncludeDeleted: parseBoolQuery(r, "include_deleted", false),
	}

	users, err := h.service.List(r.Context(), filter)
	if err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, usersResponse{Data: users})
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)
	if err != nil {
		h.handleError(w, err)
		return
	}

	user, err := h.service.GetByID(r.Context(), id, parseBoolQuery(r, "include_deleted", false))
	if err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, userResponse{Data: user})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)
	if err != nil {
		h.handleError(w, err)
		return
	}

	var input domain.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid json body"})
		return
	}

	user, err := h.service.Update(r.Context(), id, input)
	if err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, userResponse{Data: user})
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)
	if err != nil {
		h.handleError(w, err)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions

func (h *UserHandler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidUserID), errors.Is(err, domain.ErrInvalidUserInput):
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
	case errors.Is(err, domain.ErrUserNotFound):
		writeJSON(w, http.StatusNotFound, errorResponse{Error: err.Error()})
	case errors.Is(err, domain.ErrEmailAlreadyTaken):
		writeJSON(w, http.StatusConflict, errorResponse{Error: err.Error()})
	default:
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
		// writeJSON(w, http.StatusInternalServerError, errorResponse{Error: err.Error()})
	}
}

func parseIDParam(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id <= 0 {
		return 0, domain.ErrInvalidUserID
	}
	return id, nil
}

func parseIntQuery(r *http.Request, key string, fallback int) int {
	value := r.URL.Query().Get(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func parseBoolQuery(r *http.Request, key string, fallback bool) bool {
	value := r.URL.Query().Get(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}
