package user

import "testing"

func TestHashPassword(t *testing.T) {
	t.Parallel()
	password := "P@ssw0rd#123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Log("error hashing password", err)
	}

	err = CheckPassword(hashedPassword, password)
	if err != nil {
		t.Log("error comparing password", err)
	}
}
