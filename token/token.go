package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var JWT_PUBLIC_KEY *rsa.PublicKey

func init() {
	pubKeyPEM, _ := base64.StdEncoding.DecodeString(os.Getenv("JWT_PUBLIC_KEY"))
	block, _ := pem.Decode(pubKeyPEM)
	if block == nil {
		log.Fatal("Failed to parse PEM block containing the public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("Failed to parse public key: " + err.Error())
	}

	var ok bool
	JWT_PUBLIC_KEY, ok = pub.(*rsa.PublicKey)
	if !ok {
		log.Fatal("Public key of unsupported type")
	}
}

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	GoogleID    string `json:"google_id"`
	Role        string `json:"role"`
	TokenDevice string `json:"token_device"`
}

type Record struct {
	ID        int    `json:"id"`
	User      User   `json:"user"`
	Message   string `json:"message"`
	Title   	string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func DecodeToken(tokenString string) (*Record, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_PUBLIC_KEY, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		recordJSON, _ := json.Marshal(claims["record"])
		var record Record
		err := json.Unmarshal(recordJSON, &record)
		if err != nil {
			return nil, err
		}
		return &record, nil
	} else {
		return nil, err
	}
}
