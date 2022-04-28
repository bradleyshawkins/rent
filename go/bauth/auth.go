package bauth

import (
	"context"
	"log"
	"net/http"

	"github.com/bradleyshawkins/rent/berror"
	"github.com/golang-jwt/jwt/v4"
)

type authenticationKey string

const (
	authenticationContextKey authenticationKey = "authenticationKey"
)

type Authenticator interface {
	Authenticate(authentication string) (*jwt.Token, error)
}

func AddAuthenticationContext(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, authenticationContextKey, token)
}

func GetTokenFromContext(ctx context.Context) (*jwt.Token, error) {
	tokenVal := ctx.Value(authenticationContextKey)
	if tokenVal == nil {
		return nil, berror.New("authentication token not found", berror.WithUnauthenticated())
	}

	token, ok := tokenVal.(*jwt.Token)
	if !ok {
		return nil, berror.New("unexpected token value found", berror.WithUnauthenticated())
	}

	return token, nil
}

func AuthenticateMiddleware(a Authenticator) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authentication")
			token, err := a.Authenticate(authToken)
			if err != nil {
				log.Println("Authentication failed. Error:", err)
				if berr, ok := err.(*berror.Error); ok {
					berr.WriteAsJson(w)
				}
				http.Error(w, "authentication failed", http.StatusUnauthorized)
			}

			r.WithContext(AddAuthenticationContext(r.Context(), token))
		})
	}
}
