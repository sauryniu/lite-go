package cryptoex

// EncryptOption is 加密选项
type EncryptOption func(IEncryptor)

// IEncryptor is 加密接口
type IEncryptor interface {
	Encrypt(plaintext []byte, options ...EncryptOption) (ciphertext []byte, err error)
}
