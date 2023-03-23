package auth

import (
	"errors"
	"fmt"
	"go-book/method"
	"go-book/user"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(user user.User) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct{}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(user user.User) (string, error) {
	claim := jwt.MapClaims{}
	claim["user"] = user

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secretKey, err := method.GetSecretKey()

	if err != nil {
		return secretKey, err
	}

	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		var envs map[string]string
		envs, err := godotenv.Read(".env")

		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(envs["SECRET_KEY"])
		secretKey := envs["SECRET_KEY"]

		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
