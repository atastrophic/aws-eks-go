package password

import (
	"github.com/alexedwards/argon2id"
)

type PasswordGenerator struct {
	params *argon2id.Params
}

func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{
		params: &argon2id.Params{
			Memory:      128 * 1024,
			Iterations:  4,
			Parallelism: 4,
			SaltLength:  16,
			KeyLength:   32,
		},
	}
}

func (s *PasswordGenerator) Generate(plaintext string) (string, error) {
	hash, err := argon2id.CreateHash(plaintext, s.params)
	if err != nil {
		return "", err
	}
	return hash, err
}

func (s *PasswordGenerator) Compare(plaintext, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(plaintext, hash)
}
