package helper

import (
	"gzfs/Go~Edita/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Admin      bool
	jwt.RegisteredClaims
}

func GenerateAllTokens(envConfig config.Env, userEmail string, firstName string, lastName string, isAdmin bool) (signedToken string, signedRefereshToken string, err error) {
	tokenClaim := &SignedDetails{
		Email:      userEmail,
		First_Name: firstName,
		Last_Name:  lastName,
		Admin:      isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	refreshTokenClaim := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}

	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim).SignedString([]byte(envConfig.JWT_SECRET))
	if err != nil {
		log.Panic(err)
		return
	}

	signedRefereshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim).SignedString([]byte(envConfig.JWT_SECRET))

	if err != nil {
		log.Panic(err)
		return
	}

	return signedToken, signedRefereshToken, err
}
