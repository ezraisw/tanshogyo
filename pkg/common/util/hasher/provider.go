package hasher

import "golang.org/x/crypto/bcrypt"

func ProvideHasher() Hasher {
	return BcryptHasher(bcrypt.DefaultCost)
}
