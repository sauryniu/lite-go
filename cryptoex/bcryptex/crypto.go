package bcryptex

import (
	"errors"

	"github.com/ahl5esoft/lite-go/cryptoex"
	"golang.org/x/crypto/bcrypt"
)

var errDecrypt = errors.New("bcrypt不支持解密")

type crypto struct{}

func (m crypto) Decrypt(_ []byte, options ...cryptoex.DecryptOption) (plaintext []byte, err error) {
	return nil, errDecrypt
}

func (m crypto) Encrypt(plaintext []byte, _ ...cryptoex.EncryptOption) ([]byte, error) {
	return bcrypt.GenerateFromPassword(plaintext, bcrypt.DefaultCost)
}

func (m crypto) Validate(ciphertext []byte, plaintext []byte) bool {
	return bcrypt.CompareHashAndPassword(ciphertext, plaintext) == nil
}

// New is cryptoex.ICrypto实例
func New() cryptoex.ICrypto {
	return &crypto{}
}
