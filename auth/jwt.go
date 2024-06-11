package auth

/*

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofor-little/env"
)

var jwtKey = []byte(env.Get("TOKEN_SECRET", "someHiddenJwtToken"))

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   string `json:"userId"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string, userId string, expiration time.Duration) (tokenString string, err error) {
	expirationTime := time.Now().Add(expiration)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}

	return token, nil
}

func GenerateTokenPair(email string, username string, userId string) (accessToken string, refreshToken string, err error) {
	accessToken, err = GenerateJWT(email, username, userId, 1*time.Hour)
	if err != nil {
		return
	}
	refreshToken, err = GenerateJWT(email, username, userId, 2*time.Hour)
	return
}
*/
