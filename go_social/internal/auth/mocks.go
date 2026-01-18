package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TestAuthenticator struct {
}

var testSecret = "test"
var testClaims = jwt.MapClaims{
	"iat": time.Now().Unix(),
	"nbf": time.Now().Unix(),
	"aud": "test-aud",
	"iss": "test-iss",
	"sub": int(42),
	"exp": time.Now().Add(time.Hour).Unix(),
}

func (a *TestAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, testClaims)

	tokenString, _ := token.SignedString([]byte(testSecret))

	return tokenString, nil
}

func (a *TestAuthenticator) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (any, error) {

		return []byte(testSecret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience("test-aud"),
		jwt.WithIssuer("test-iss"),
	)
}

func NewMockAuth() *TestAuthenticator {
	return &TestAuthenticator{}
}
