package security

import (
	"crypto/aes"
	"crypto/sha256"
	"playdates/internal/secretmanager"
)

var (
	encryptionSecretName = "dev_encryption_key_base64"
	hashingSecretName    = "dev_hash_key_base64"
)

type Security struct {
	encryptSecret []byte
	hashSecret    []byte
}

func NewSecurity(secretManager *secretmanager.SecretManagerClient, projectID string) (*Security, error) {
	encryptSecret, err := secretManager.GetSecret(encryptionSecretName)
	if err != nil {
		return nil, err
	}

	hashSecret, err := secretManager.GetSecret(hashingSecretName)
	if err != nil {
		return nil, err
	}

	return &Security{encryptSecret: encryptSecret, hashSecret: hashSecret}, nil
}

func (s *Security) Encrypt(data []byte) ([]byte, error) {
	// Create a new AES encrypter
	encrypter, err := aes.NewCipher(s.encryptSecret)
	if err != nil {
		return nil, err
	}

	// Encrypt
	encrypted := make([]byte, len(data))
	encrypter.Encrypt(encrypted, data)

	return encrypted, nil
}

func (s *Security) Hash(data []byte) ([]byte, error) {
	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Hash the data
	hasher.Write(data)
	hash := hasher.Sum(nil)

	return hash, nil
}
