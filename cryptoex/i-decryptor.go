package cryptoex

// DecryptOption is 解密选项
type DecryptOption func(IDecryptor)

// IDecryptor is 解密接口
type IDecryptor interface {
	Decrypt(ciphertext []byte, options ...DecryptOption) (plaintext []byte, err error)
}
