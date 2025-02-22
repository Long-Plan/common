package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type jwtSigner struct {
	privateKey []byte
	publicKey  []byte
}

func NewJWTSigner(privateKey []byte, publicKey []byte) IJWTSigner {
	return &jwtSigner{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

type IJWTSigner interface {
	SignToken(token *jwt.Token) (string, error)
	NewClaims(claims jwt.Claims) *jwt.Token
}

func (j *jwtSigner) SignToken(token *jwt.Token) (string, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", err
	}
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *jwtSigner) NewClaims(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
}
