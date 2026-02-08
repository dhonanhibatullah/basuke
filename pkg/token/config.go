package token

import "time"

var (
	issuer     string
	key        []byte
	expiration time.Duration
)

func Config(issuerName string, privateKey string, expireDuration time.Duration) {
	issuer = issuerName
	key = []byte(privateKey)
	expiration = expireDuration
}
