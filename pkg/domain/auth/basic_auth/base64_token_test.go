package basic_auth

import (
	"encoding/base64"
	"testing"
)

func TestDecodeBase64Token(t *testing.T) {

	var a base64.CorruptInputError = 8
	corruptInputErr := a.Error()
	type args struct {
		token string
	}
	testsCases := []struct {
		name         string
		args    args
		want    Base64Token
		wantErr bool
		errorMessage string
	}{
		{
			name: "should decode successfully",
			args: args{
				token: "bWV1OnNlZ3JlZG8=",
			},
			want: Base64Token{"meu", "segredo"},
		},

		{
			name: "should get err on decode",
			args: args{
				token: "something",
			},
			wantErr:      true,
			errorMessage: corruptInputErr,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := DecodeBase64Token(tc.args.token)

			if tc.wantErr {
				if tc.errorMessage != err.Error() {
					t.Fatalf("DecodeBase64Token error message. want %s got %s", tc.want.APIKey, got.APIKey)
				}
				return
			}

			if err != nil {
				t.Fatal("unexpected error", err)
			}

			if got.APIKey != tc.want.APIKey {
				t.Fatalf("DecodeBase64Token apiKey. want %s got %s", tc.want.APIKey, got.APIKey)
			}
			if got.Secret != tc.want.Secret {
				t.Fatalf("DecodeBase64Token secret. want %s got %s", tc.want.Secret, got.Secret)
			}
		})
	}
}
