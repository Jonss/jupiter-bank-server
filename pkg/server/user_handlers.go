package server

import (
	"encoding/json"
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
			apiResponse(w, http.StatusBadRequest, nil)
			return
		}

		err := s.restValidator.Validator.Struct(req)
		if err != nil {
			ValidateRequestBody(err, w, s.restValidator.Translator)
			return
		}

		u, err := s.userService.CreateUser(ctx, user.CreateUseParams{
			Fullname: req.Fullname,
			Email:    req.Email,
			Password: req.Password,
		})

		if err != nil {
			if err == user.ErrUserExists {
				apiResponse(w, http.StatusUnprocessableEntity, UserExists)
				return
			}
			apiResponse(w, http.StatusInternalServerError, UnexpectedError)
			return
		}

		apiResponse(w, http.StatusCreated, createUserResponse{
			ExternalID: u.ExternalID,
			CreatedAt:  u.CreatedAt,
		})
	}
}

func (s *Server) Profile() http.HandlerFunc {
	type response struct {
		ExternalID uuid.UUID `json:"external_id"`
		Fullname   string    `json:"fullname"`
		Email      string    `json:"email"`
		CreatedAt  time.Time `json:"created_at"`
		TaxID      *string   `json:"tax_id"`
		IsComplete bool      `json:"is_complete"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		userID := UserIDFromRequest(r)

		u, err := s.userService.GetUserByID(r.Context(), userID)
		if err != nil {
			apiResponse(w, http.StatusInternalServerError, UnexpectedError)
			return
		}
		resp := response{
			ExternalID: u.ExternalID,
			Fullname:   u.Fullname,
			Email:      u.Email,
			CreatedAt:  u.CreatedAt,
			TaxID:      u.TaxID,
			IsComplete: u.IsComplete(),
		}
		apiResponse(w, http.StatusOK, resp)
	}

}
