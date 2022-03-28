package paseto_auth

import (
	"errors"
	"testing"
)

func TestCreateTokenAndDecryptToken(t *testing.T) {
	successToken, err := CreateToken(1)
	if err != nil {
		t.Fatal("error when create token", err)
	}

	testCases := []struct {
		name       string
		signedKey  string
		publicHex  string
		wantError  error
		wantUserID string
	}{
		{
			name:       "should decrypt token",
			publicHex:  successToken.PublicHex,
			signedKey:  successToken.SignedKey,
			wantError:  nil,
			wantUserID: "1",
		},
		{
			name:       "should get an error when key is expired",
			publicHex:  "1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2",
			signedKey:  "v4.public.eyJkYXRhIjoidGhpcyBpcyBhIHNpZ25lZCBtZXNzYWdlIiwiZXhwIjoiMjAyMi0wMS0wMVQwMDowMDowMCswMDowMCJ9v3Jt8mx_TdM2ceTGoqwrh4yDFn0XsHvvV_D0DtwQxVrJEBMl0F2caAdgnpKlt4p7xBnx1HcO-SPo8FPp214HDw.eyJraWQiOiJ6VmhNaVBCUDlmUmYyc25FY1Q3Z0ZUaW9lQTlDT2NOeTlEZmdMMVc2MGhhTiJ9",
			wantError:  errors.New("this token has expired"), // TODO - assert the error from paseto lib
			wantUserID: "90",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userID, err := DecryptToken(tc.signedKey, tc.publicHex)
			if tc.wantError != nil && tc.wantError.Error() != err.Error() {
				t.Fatal("error when decrypt token.", err)
			}
			if tc.wantError == nil {
				if userID != tc.wantUserID {
					t.Fatalf("want %s got  %s", tc.wantUserID, userID)
				}
			}
		})
	}

}
