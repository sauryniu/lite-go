package cryptoex

// DecryptOption is 解密选项
type DecryptOption func(ICrypto)

// EncryptOption is 加密选项
type EncryptOption func(ICrypto)

// ICrypto is 加密接口
type ICrypto interface {
	Decrypt(ciphertext []byte, options ...DecryptOption) (plaintext []byte, err error)
	Encrypt(plaintext []byte, options ...EncryptOption) (ciphertext []byte, err error)
	Validate(ciphertext []byte, plaintext []byte) bool
}
