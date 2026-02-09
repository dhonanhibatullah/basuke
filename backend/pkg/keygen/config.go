package keygen

var (
	secretKey []byte = []byte("secret")
)

func Config(secret string) {
	secretKey = []byte(secret)
}
