package jwtservice

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var SecretKey = []byte("secret-key")

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func CreateToken(userID uuid.UUID, email string) (Tokens, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"email":   email,
			"exp":     time.Now().Add(time.Minute * 15).Unix(),
		})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return Tokens{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"email":   email,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})
	refreshTokenString, err := refreshToken.SignedString(SecretKey)
	if err != nil {
		return Tokens{}, nil
	}

	t := Tokens{
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}

	return t, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("cannot extract claims")
	}

	return claims, nil
}

// type Claims struct {
// 	UserID string `json:"uid"`
// 	Email  string `json:"email"`
// 	jwt.RegisteredClaims
// }

// func CreateToken(userID uuid.UUID, email string, secret []byte, expiresIn time.Duration) (string, error) {
// 	claims := &Claims{
// 		UserID: userID.String(),
// 		Email:  email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(secret)
// }

// func ParseToken(tokenString string, secret []byte) (*Claims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return secret, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
// 		return claims, nil
// 	}

// 	return nil, fmt.Errorf("invalid token")
// }

//

// package jwt

// import (
// 	"fmt"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// var secretKey = []byte("the_most_secret_key")

// func CreateToken(email string) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodES256,
// 		jwt.MapClaims{
// 			"email": email,
// 			"exp":   time.Now().Add(time.Hour * 24).Unix(),
// 		})

// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// func VerifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if !token.Valid {
// 		return fmt.Errorf("invalid token")
// 	}

// 	return nil
// }
