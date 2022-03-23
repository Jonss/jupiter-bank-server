package server

import (
	"encoding/json"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"net/http"
	"time"

	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/google/uuid"
)

type createUserRequest struct {
	Fullname string `json:"fullname" validate:"required,gte=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6,lte=100"`
}

type createUserResponse struct {
	ExternalID uuid.UUID `json:"external_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (s *Server) Signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req createUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			rest.JsonResponse(w, http.StatusBadRequest, nil)
			return
		}

		u, err := s.userService.CreateUser(ctx, user.CreateUseParams{
			Fullname: req.Fullname,
			Email:    req.Email,
			Password: req.Password,
		})

		if err != nil {
			if err == user.ErrUserExists {
				rest.JsonResponse(w, http.StatusUnprocessableEntity, rest.UserExists)
				return
			}
			rest.JsonResponse(w, http.StatusInternalServerError, rest.UnexpectedError)
			return
		}

		rest.JsonResponse(w, http.StatusCreated, createUserResponse{
			ExternalID: u.ExternalID,
			CreatedAt:  u.CreatedAt,
		})
	}
}
