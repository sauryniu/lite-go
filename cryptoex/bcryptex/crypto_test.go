package bcryptex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_crypto_Decrypt(t *testing.T) {
	res, err := crypto{}.Decrypt(nil)
	assert.Error(t, err)
	assert.Equal(t, err, errDecrypt)
	assert.Nil(t, res)
}

func Test_crypto_Encrypt(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		plaintext := "123456"
		res, err := crypto{}.Encrypt(
			[]byte(plaintext),
		)
		assert.NoError(t, err)

		err = bcrypt.CompareHashAndPassword(
			res,
			[]byte(plaintext),
		)
		assert.NoError(t, err)
	})

	t.Run("diff", func(t *testing.T) {
		plaintext := "123456"
		res, err := crypto{}.Encrypt(
			[]byte(plaintext),
		)
		assert.NoError(t, err)

		err = bcrypt.CompareHashAndPassword(
			res,
			[]byte("0123456"),
		)
		assert.Error(t, err)
	})
}

func Test_crypto_Validate(t *testing.T) {
	plaintext := "123456"
	ciphertext, err := bcrypt.GenerateFromPassword(
		[]byte(plaintext),
		bcrypt.DefaultCost,
	)
	assert.NoError(t, err)

	res := crypto{}.Validate(
		ciphertext,
		[]byte(plaintext),
	)
	assert.True(t, res)
}
