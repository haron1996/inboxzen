package paseto

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/haron1996/inboxzen/viper"
)

func CreateToken(userID, email string, issuedAt time.Time, duration time.Duration) (string, *PayLoad, error) {
	expiry := time.Now().UTC().Add(duration)

	payload := newPayload(userID, email, issuedAt, expiry)

	token := paseto.NewToken()

	token.SetExpiration(expiry)

	token.SetIssuedAt(issuedAt)

	token.Set("payload", payload)

	config, err := viper.LoadConfig(".")
	if err != nil {
		return "", nil, err
	}

	secretKey, _ := paseto.NewV4AsymmetricSecretKeyFromHex(config.SecretKeyHex)

	signed := token.V4Sign(secretKey, nil)

	return signed, payload, nil
}
