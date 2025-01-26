package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Claims defines the data stored in the JWT token
type Claims struct {
	UserID    string    `json:"userID"`
	UserName  string    `json:"userName"`
	ExpiresAt time.Time `json:"expiry"`
	jwt.RegisteredClaims
}

type Tokens struct {
	AccessToken string `json:"accessToken"`
}

func ValidateToken(accesstokenString string) (*Claims, error) {

	claims, err := GetTokenClaims(accesstokenString)

	if err != nil {
		return nil, fmt.Errorf("failed to parse claims")
	}

	accessTokenExpiry := claims.ExpiresAt

	if accessTokenExpiry.Before(time.Now()) {
		log.Printf(" Access token expiredddd")
		return nil, fmt.Errorf("token Expired")

	}

	return claims, nil

}

func GetTokenClaims(accesstokenString string) (*Claims, error) {
	//parsing acess tokens
	accessToken, err := jwt.ParseWithClaims(accesstokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := accessToken.Claims.(*Claims)

	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	return claims, nil
}

func GenerateAcessToken(payload *Claims) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	acessToken, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", NewErrorStruct(400, "error generating JWT Token")
	}

	return acessToken, nil

}

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", NewErrorStruct(400, "bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", NewErrorStruct(400, "incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
