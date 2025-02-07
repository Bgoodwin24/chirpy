package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	password3 := "CoRrEcTpAsSwOrD789!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)
	hash3, _ := HashPassword(password3)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
		{
			name:     "Correct password2",
			password: password3,
			hash:     hash3,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func TestCreateValidateJWT(t *testing.T) {
	secret1 := "secret"
	secret2 := "wrongSecret"
	userID := uuid.New()
	token1, _ := MakeJWT(userID, secret1, time.Hour)
	token2, _ := MakeJWT(userID, secret2, time.Hour)
	expiredToken, _ := MakeJWT(userID, secret2, -time.Second)

	tests := []struct {
		name           string
		token          string
		secret         string
		wantErr        bool
		expectedUserID uuid.UUID
	}{
		{
			name:           "correct token",
			token:          token1,
			secret:         secret1,
			wantErr:        false,
			expectedUserID: userID,
		},
		{
			name:           "incorrect token",
			token:          token2,
			secret:         secret1,
			wantErr:        true,
			expectedUserID: uuid.Nil,
		},
		{
			name:           "incorrect secret",
			token:          token1,
			secret:         secret2,
			wantErr:        true,
			expectedUserID: uuid.Nil,
		},
		{
			name:           "empty secret",
			token:          token2,
			secret:         "",
			wantErr:        true,
			expectedUserID: uuid.Nil,
		},
		{
			name:           "expired token",
			token:          expiredToken,
			secret:         secret2,
			wantErr:        true,
			expectedUserID: uuid.Nil,
		},
		{
			name:           "malformed token",
			token:          "invalid.token.here",
			secret:         secret1,
			wantErr:        true,
			expectedUserID: uuid.Nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userID, err := ValidateJWT(tt.token, tt.secret)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && userID != tt.expectedUserID {
				t.Errorf("ValidateJWT() returned wrong userID = %v, expected %v", userID, tt.expectedUserID)
			}
		})
	}
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "valid bearer token",
			headers: http.Header{"Authorization": []string{"Bearer abc123"}},
			want:    "abc123",
			wantErr: false,
		},
		{
			name:    "missing bearer prefix",
			headers: http.Header{"Authorization": []string{"abc123"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty token",
			headers: http.Header{"Authorization": []string{"Bearer "}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "missing authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBearerToken(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBearerToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
