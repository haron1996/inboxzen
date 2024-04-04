package paseto

import (
	"errors"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/haron1996/inboxzen/viper"
)

func VerifyToken(signedToken string) (*PayLoad, error) {
	config, err := viper.LoadConfig(".")
	if err != nil {
		return nil, err
	}

	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(config.PublicKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error constructing a v4 public key from hex: %v", err)
	}

	parser := paseto.NewParser()

	token, err := parser.ParseV4Public(publicKey, signedToken, nil)
	if err != nil {
		return nil, fmt.Errorf("error parsing, verifying and validation v4 public: %v", err)
	}

	var payload PayLoad

	if err := token.Get("payload", &payload); err != nil {
		return nil, fmt.Errorf("error getting payload: %v", err)
	}

	if time.Now().UTC().After(payload.Expiry) {
		return nil, errors.New("token is expired")
	}

	return &payload, nil
}
