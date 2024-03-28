package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
)

const (
	jwtTTL     = time.Minute * 15
	refreshTTL = time.Hour * 24 * 30
)

type TokenService struct {
	secret []byte
}

func NewTokenService(signSecret []byte) *TokenService {
	return &TokenService{
		secret: signSecret,
	}
}

// Issue tokens to an account
//
// Returns *data.Tokens, which contains signed JWT and Refresh tokens
//
// Returns an error if something went wrong with signing jwt or generating refresh
func (ts *TokenService) IssueTokens(
	account *data.Account,
) (*data.Tokens, error) {
	tokenJWT := ts.generateJWT(account.ID.String())
	tokenJWTSigned, err := ts.signJWT(tokenJWT)
	if err != nil {
		return nil, err
	}

	tokenRefresh, err := ts.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &data.Tokens{
		JWTToken:     tokenJWTSigned,
		RefreshToken: tokenRefresh,
	}, nil
}

func (ts *TokenService) generateJWT(subject string) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(jwtTTL)},
		Subject:   subject,
	})

	return token
}

func (ts *TokenService) signJWT(token *jwt.Token) (string, error) {
	tokenStr, err := token.SignedString(ts.secret)
	if err != nil {
		return "", nil
	}

	return tokenStr, nil
}

func (ts *TokenService) generateRefreshToken() (string, error) {
	refreshToken := "amogus"

	return refreshToken, nil
}
