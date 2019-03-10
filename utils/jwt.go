package utils

import (
	"auxpi/bootstrap"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(bootstrap.SiteConfig.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	//V        uint   `json:"v"`
	jwt.StandardClaims
}

func GenerateToken(username, email string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add( bootstrap.SiteConfig.JwtDueTime * time.Hour)

	claims := Claims{
		username,
		email,
		//version,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    bootstrap.SiteConfig.SiteName,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
