package jwt

import (
	"logoinig/internal/services/database/postgredb"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_Secret_Key")

type (
	LoginInfo struct {
		Password string
		Username string
	}
	Claims struct {
		Username string
		jwt.StandardClaims
	}
)

func SingIn(login LoginInfo) (stringToken string, ok bool) {
	user, err := postgredb.Hand.ReadUser(login.Username)
	if err != nil {
		return
	}
	if login.Password != user.Password {
		return
	}
	expTime := time.Now().Add(5 * time.Minute)
	cliams := &Claims{
		Username: login.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	stringToken, err = token.SignedString(jwtKey)
	if err != nil {
		return
	}
	ok = true
	// fmt.Println("the cookie: ", stringToken)
	return
}

func IsLogedin(stringToken string) (lg LoginInfo, ok bool) {
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(stringToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return
	}
	lg.Username, ok = claims.Username, true
	return
}
