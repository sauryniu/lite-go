package cryptoex

// IValidator is 验证器
type IValidator interface {
	Validate(ciphertext []byte, plaintext []byte) bool
}
