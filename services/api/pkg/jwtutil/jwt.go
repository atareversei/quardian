package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var globalJWT *JWT

type JWTConfig struct {
	// Expiration time in days.
	BasicTokenExpirationTime int    `koanf:"basic_token_expiration_time_days"`
	BasicTokenSecretKey      string `koanf:"basic_token_secret_key"`
}

type JWT struct {
	config JWTConfig
}

func Init(cfg JWTConfig) {
	jwt := &JWT{
		config: cfg,
	}

	globalJWT = jwt
}

func Create(userId int) (string, error) {
	basic, err := CreateBasicToken(userId)
	if err != nil {
		return "", err
	}

	return basic, nil
}

func CreateBasicToken(userId int) (string, error) {
	expTime := time.Hour * 24 * time.Duration(globalJWT.config.BasicTokenExpirationTime)

	claims := jwt.MapClaims{
		"user_id": userId,
		"iss":     "api.quardian",
		"exp":     time.Now().Add(expTime).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(globalJWT.config.BasicTokenSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// TODO: rewrite this function
func Decode(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("")
		}
		return []byte(globalJWT.config.BasicTokenSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("")
	}
}

// TODO: error handling
func ValidateAndGetUserId(tokenString string) (int, error) {
	claims, err := Decode(tokenString)
	if err != nil {
		return 0, err
	}

	// TODO: redundant check?
	c, ok := claims["user_id"]
	if !ok {
		return 0, err
	}

	value, ok := c.(float64)

	if !ok {
		// TODO: enhance error
		return 0, nil
	}

	return int(value), err
}
