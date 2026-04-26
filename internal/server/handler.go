package server

import (
	"context"

	"github.com/Bayan2019/rbk-it-school-hw-3/internal/domain"

	httptransport "github.com/Bayan2019/rbk-it-school-hw-3/internal/transport/http"
)

type userService interface {
	Create(ctx context.Context, input domain.CreateUserInput) (domain.User, error)
	List(ctx context.Context, filter domain.ListUsersFilter) ([]domain.User, error)
	GetByID(ctx context.Context, id int64, includeDeleted bool) (domain.User, error)
	Update(ctx context.Context, id int64, input domain.UpdateUserInput) (domain.User, error)
	Delete(ctx context.Context, id int64) error
}

type Handler struct {
	User *httptransport.UserHandler
}

func NewHandler(userService userService) *Handler {
	return &Handler{
		User: httptransport.NewUserHandler(userService),
	}
}
