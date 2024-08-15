package src

import (
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil && len(bytes) != 0 {
		return "", err
	}

	hashed := string(bytes)

	return hashed, err
}

func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var PayloadData = new(Payload)

type Payload struct {
	User User
}

var SecretKey = "custom-dating-private-key"

func GenerateToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.Base.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	bearerToken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return "Bearer " + bearerToken, nil

}

func JwtKeyFunc() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwtware.HS256 {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		signingKey := "custom-dating-private-key"
		return []byte(signingKey), nil
	}
}

func ParsePayload(ctx *fiber.Ctx) {
	jwtToken := ctx.Locals("jwtKey").(*jwt.Token)
	claims := *(jwtToken.Claims.(*jwt.MapClaims))

	id := int(claims["id"].(float64))
	PayloadData = &Payload{
		User{
			Base:     Base{ID: id},
			Email:    claims["email"].(string),
			Username: claims["username"].(string),
		},
	}
}
