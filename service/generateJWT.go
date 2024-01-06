package service 


import (
	"time"

	"github.com/dgrijalva/jwt-go"
)


func GenerateJWTToken(username, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "anurag",
		Subject:   username,
		Audience:  role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("e-comm-secret-key")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
