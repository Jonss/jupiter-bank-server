package server

import (
	"encoding/json"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
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
			rest.JsonResponse(w, http.StatusBadRequest, nil)
			return
		}

		err := s.restValidator.Validator.Struct(req)
		if err != nil {
			rest.ValidateRequestBody(err, w, s.restValidator.Translator)
			return
		}
		pasetoToken, err := s.pasetoAuthService.Login(ctx, req.Email, req.Password)
		if err != nil {
			rest.JsonResponse(w, http.StatusInternalServerError, rest.NewErrorResponses(rest.ErrorResponse{
				Code:    rest.CodeAuth1,
				Message: err.Error(),
			}))
			return
		}

		resp := response{
			Hex:   pasetoToken.PublicHex,
			Token: pasetoToken.SignedKey,
		}
		rest.JsonResponse(w, http.StatusOK, resp)
	}
}
