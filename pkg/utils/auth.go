package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type customClaims struct {
	Id        uint16 `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.StandardClaims
}

var SECURE_SECRET_TEXT string = "PLnMAvreoND23DkAedMj0Wx501Rgt8lPwJBVjtM2YKsyw1fEH7XCI94N6R8UDlAxY7Vip3oWSg2dfbtQGMTnuSelbxG0RIyel0k4kniKJmAMLKYY7xpf0QoXOBfhvVV5UK0f2L0tvOBGt8VYs2PFJJkw4mxi9Va9qoUar7QxfLvJQclmtRAjR6CVo1qimaEeU7jBeOCeKs6ckFvXFBpD4sLzlFyY8JKKQR4k581bdMlRN17jr66tet96iemZKqoG"

func GenerateJWT(id uint16, email string, firstName string, lastName string) (string, error) {
	claims := customClaims{
		Id:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(SECURE_SECRET_TEXT))

	return signedToken, err
}

func ValidateJWT(jwtToken string) (customClaims, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECURE_SECRET_TEXT), nil
		},
	)

	if err != nil {
		return customClaims{}, err
	}

	claims, ok := token.Claims.(*customClaims)

	if !ok {
		return customClaims{}, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return customClaims{}, errors.New("jwt is expired")
	}

	return *claims, nil
}
