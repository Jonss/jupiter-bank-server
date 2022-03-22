package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type createUserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	ExternalID uuid.UUID `json:"external_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (d *UserDomain) Signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req createUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			// TODO: group return error in a individual function
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user, err := d.createUser(ctx, req)
		if err != nil {
			fmt.Println(err)
			// TODO: group return error in a individual function
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := createUserResponse{
			ExternalID: user.ExternalID,
			CreatedAt:  user.CreatedAt,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
