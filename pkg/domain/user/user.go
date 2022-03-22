package user

import "github.com/Jonss/jupiter-bank-server/pkg/db"

type UserDomain struct {
	q *db.Queries
}

func NewUserDomain(q *db.Queries) *UserDomain {
	return &UserDomain{q: q}
}
