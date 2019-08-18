package auth

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

const (
	JWT_SIGNING_KEY_TYPE = "JWT_SIGNING_KEY_TYPE"
	HS256                = "HS256"
	RS256                = "RS256"

	HS256_SIGNING_KEY         = "HS256_SIGNING_KEY"
	RS256_PRIVATE_SECRET_PATH = "RS256_PRIVATE_SECRET_PATH"
	RS256_PUBLIC_PATH         = "RS256_PUBLIC_PATH"
)

func GetJwtTokenHandler(w http.ResponseWriter, r *http.Request) {
	name := "trewanek"
	token := ""
	if os.Getenv(JWT_SIGNING_KEY_TYPE) == HS256 {
		token, _ = createJwtTokenFromHS256Secret(name)
	}
	if os.Getenv(JWT_SIGNING_KEY_TYPE) == RS256 {
		token, _ = createJwtTokenFromRS256PEM(name)
	}

	if token == "" {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("token string could not created."))
		return
	}
	_, _ = w.Write([]byte(token))
}

func JwtMiddleware(handler http.Handler) http.Handler {
	if os.Getenv(JWT_SIGNING_KEY_TYPE) == HS256 {
		return jwtMiddlewareHS256(handler)
	}
	if os.Getenv(JWT_SIGNING_KEY_TYPE) == RS256 {
		return jwtMiddlewareRS256(handler)
	}
	panic("jwt auth method is not selected.")
}

func createJwtTokenFromHS256Secret(userName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	setClaims(token, userName)
	signKey := []byte(os.Getenv(HS256_SIGNING_KEY))
	return token.SignedString(signKey)
}

func createJwtTokenFromRS256PEM(userName string) (string, error) {
	signBytes, err := ioutil.ReadFile(os.Getenv(RS256_PRIVATE_SECRET_PATH))
	if err != nil {
		return "", err
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodRS256)
	setClaims(token, userName)
	return token.SignedString(signKey)
}

func setClaims(token *jwt.Token, userName string) jwt.MapClaims {
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = userName
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return claims
}

func jwtMiddlewareHS256(handler http.Handler) http.Handler {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv(HS256_SIGNING_KEY)), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	}).Handler(handler)
}

func jwtMiddlewareRS256(handler http.Handler) http.Handler {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			b, err := ioutil.ReadFile(os.Getenv(RS256_PUBLIC_PATH))
			if err != nil {
				return nil, err
			}
			signKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
			if err != nil {
				return nil, errors.New("private key could not be parsed")
			}
			return signKey, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	}).Handler(handler)
}
