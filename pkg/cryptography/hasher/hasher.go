package hasher

import (
	"crypto/sha512"

	"github.com/pkg/errors"
	"github.com/smotes/pwdhash"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
	Cost                   = 1000
)

func GenerateFromPassword(password string) (encodedHash string, err error) {
	salt, err := pwdhash.GenerateSalt(sha512.Size)
	if err != nil {
		panic(err)
	}
	hpwd, err := pwdhash.GenerateFromPassword([]byte(password), salt, Cost, sha512.Size, "sha512")
	return string(hpwd), nil
}

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	result := pwdhash.CompareHashAndPassword([]byte(encodedHash), []byte(password))
	if result != nil {
		return false, result
	} else {
		return true, nil
	}
}
