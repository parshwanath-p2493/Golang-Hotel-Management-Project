package helpers

import (
	//	"github.com/dgrijalva/jwt-go"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Info struct {
	Name       string
	Email      string
	Role       string
	Department string
	jwt.StandardClaims
}

//var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateToken(name string, email string, role string, dept string) (string, error) {
	claims := &Info{
		Name:       name,
		Email:      email,
		Role:       role,
		Department: dept,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")
	fmt.Println("Secret key", SECRET_KEY)
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //.SignedString([]byte(SECRET_KEY))
	fmt.Printf("SECRET_KEY: %s", SECRET_KEY)

	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Print("Error signing the token", err)
		return " --->>", err
	}
	fmt.Printf("signedstring: %s", signedToken)
	return signedToken, nil
}
