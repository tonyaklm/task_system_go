package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"task_system_go/config"
	"task_system_go/models"
	"time"
)

type SignedDetails struct {
	Username string
	UserID   uint
	jwt.StandardClaims
}

func GenerateAllTokens(user models.User) (string, string, error) {
	claims := &SignedDetails{
		Username: user.Username,
		UserID:   user.ID,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(config.Cfg.Server.ExpirationMinutes)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(config.Cfg.Server.ExpirationHours)).Unix(),
		},
	}

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(config.Cfg.Server.SecretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).
		SignedString([]byte(config.Cfg.Server.SecretKey))
	if err != nil {
		return "", "", err
	}

	return signedToken, refreshToken, nil
}

func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.Server.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		err = errors.New("the token is invalid")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return
	}
	return
}
