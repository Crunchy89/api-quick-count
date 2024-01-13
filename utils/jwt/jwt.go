// cSpell:ignore issure
package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// jwt service
type JWTHelper interface {
	GenerateToken(UUID uuid.UUID, tokenAccount string) string
	ValidateToken(token string) (*jwt.Token, error)
	GenerateTokenPublic(UUID uuid.UUID) string
}
type authCustomClaims struct {
	UUID  uuid.UUID `json:"uuid"`
	Token string    `json:"token"`
	jwt.StandardClaims
}
type authCustomPublicClaims struct {
	UUID uuid.UUID `json:"uuid"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

// auth-jwt
func NewJWTService() JWTHelper {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "api-techcode",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET_KEY")
	return secret
}

func (service *jwtServices) GenerateToken(UUID uuid.UUID, tokenAccount string) string {
	now := time.Now()
	claims := &authCustomClaims{
		UUID,
		tokenAccount,
		jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (service *jwtServices) GenerateTokenPublic(UUID uuid.UUID) string {
	now := time.Now()
	claims := &authCustomPublicClaims{
		UUID,
		jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

}
