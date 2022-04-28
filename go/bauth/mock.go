package bauth

import (
	"github.com/golang-jwt/jwt/v4"
)

type Mock struct {
	AuthenticationParam string
	AuthenticateToken   *jwt.Token
	AuthenticateError   error
}

func (m *Mock) Authenticate(authentication string) (*jwt.Token, error) {
	m.AuthenticationParam = authentication
	return m.AuthenticateToken, m.AuthenticateError
}
