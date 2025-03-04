package helpers

import (
	//	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Info struct {
	Name  string
	Email string
	Role  string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateToken(name string, email string, role string) (string, error) {
	claims := &Info{
		Name:  name,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //.SignedString([]byte(SECRET_KEY))
	// Sign the token with your secret key
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Print("Error signing the token", err)
		return "", err
	}
	return signedToken, nil
}
