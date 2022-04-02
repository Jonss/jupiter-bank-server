package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (s *Server) Authenticate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	type response struct {
		Token string `json:"token"`
		Hex   string `json:"hex"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			apiResponse(w, http.StatusBadRequest, nil)
			return
		}

		err := s.restValidator.Validator.Struct(req)
		if err != nil {
			ValidateRequestBody(err, w, s.restValidator.Translator)
			return
		}
		pasetoToken, err := s.pasetoAuthService.Login(ctx, req.Email, req.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				apiResponse(w, http.StatusNotFound, NewErrorResponses(ErrorResponse{
					Code:    CodeUser3,
					Message: err.Error(),
				}))
			}

			apiResponse(w, http.StatusInternalServerError, NewErrorResponses(ErrorResponse{
				Code:    CodeAuth1,
				Message: err.Error(),
			}))
			return
		}

		resp := response{
			Hex:   pasetoToken.PublicHex,
			Token: pasetoToken.SignedKey,
		}
		apiResponse(w, http.StatusOK, resp)
	}
}
