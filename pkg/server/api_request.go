package server

import "net/http"

const userIDKey = "userID"

func UserIDFromRequest(r *http.Request) uint64 {
	return r.Context().Value(userIDKey).(uint64)
}
