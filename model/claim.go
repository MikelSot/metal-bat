package model

import (
	"crypto/rsa"
	"sync"

	"github.com/golang-jwt/jwt"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

type Claim struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	IPClient string `json:"ip_client"`
	UserType uint   `json:"user_type"`
	jwt.StandardClaims
}

func LoadSignatures(private, public []byte, logger Logger) {
	once.Do(func() {
		var err error
		signKey, err = jwt.ParseRSAPrivateKeyFromPEM(private)
		if err != nil {
			logger.Fatalf("authorization.LoadSignatures: realizando el parse en jwt RSA private: %s", err)
		}

		verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(public)
		if err != nil {
			logger.Fatalf("authorization.LoadSignatures: realizando el parse en jwt RSA public: %s", err)
		}
	})
}

func SignKey() *rsa.PrivateKey {
	return signKey
}

func VerifyKey() *rsa.PublicKey {
	return verifyKey
}
