package bauth

import (
	"crypto/rsa"
	"strings"

	"github.com/bradleyshawkins/rent/berror"
	"github.com/golang-jwt/jwt/v4"
)

type JWTAuthenticator struct {
	pubKey *rsa.PublicKey
	parser *jwt.Parser
}

func NewJWTAuthenticator(publicKey []byte) (*JWTAuthenticator, error) {

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, berror.WrapInternal(err, "unable to parse public key")
	}
	p := jwt.NewParser(jwt.WithValidMethods([]string{
		jwt.SigningMethodRS256.Alg(),
	}))
	return &JWTAuthenticator{
		pubKey: key,
		parser: p,
	}, nil
}

func (j *JWTAuthenticator) Authenticate(authentication string) (*jwt.Token, error) {
	if len(authentication) == 0 {
		return nil, berror.New("authentication header not provided", berror.WithAuthenticationFailed())
	}

	authParts := strings.Split(authentication, ` `)
	if len(authParts) != 2 {
		return nil, berror.New("invalid authentication header provided", berror.WithAuthenticationFailed())
	}

	if authParts[0] != "Bearer" {
		return nil, berror.New("non bearer authentication header provided", berror.WithAuthenticationFailed())
	}

	token, err := j.parser.Parse(authParts[1], j.keyFunc)
	if err != nil {
		return nil, berror.Wrap(err, berror.WithMessage("invalid authentication header provided"), berror.WithAuthenticationFailed())
	}

	return token, nil
}

func (j *JWTAuthenticator) keyFunc(token *jwt.Token) (interface{}, error) {
	return j.pubKey, nil
}
