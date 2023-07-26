package hasher

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type (
	Hasher interface {
		Hash(str string) (string, error)
		Compare(str string, hashedStr string) (bool, error)
	}

	bcryptHasher struct {
		cost int
	}

	noHasher struct{}
)

func BcryptHasher(cost int) Hasher {
	return &bcryptHasher{cost: cost}
}

func (h bcryptHasher) Hash(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), h.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (bcryptHasher) Compare(str string, hashedStr string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(str))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

var NoHasher Hasher = &noHasher{}

func (noHasher) Hash(str string) (string, error) {
	return str, nil
}

func (noHasher) Compare(str string, hashedStr string) (bool, error) {
	return str == hashedStr, nil
}
