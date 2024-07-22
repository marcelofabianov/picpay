package hasher

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Argon2Hasher struct{}

func NewArgon2Hasher() *Argon2Hasher {
	return &Argon2Hasher{}
}

func (h *Argon2Hasher) Hash(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=19$m=65536,t=1,p=4$%s$%s", b64Salt, b64Hash)

	return encodedHash, nil
}

func (h *Argon2Hasher) Compare(hashedPassword string, password string) error {
	var salt, hash []byte
	var err error

	parts := make([][]byte, 3)

	for i, part := range parts {
		parts[i] = []byte(part)
	}

	_, err = fmt.Sscanf(hashedPassword, "$argon2id$v=19$m=65536,t=1,p=4$%s$%s", &parts[0], &parts[1])
	if err != nil {
		return err
	}

	salt, err = base64.RawStdEncoding.DecodeString(string(parts[0]))
	if err != nil {
		return err
	}

	hash, err = base64.RawStdEncoding.DecodeString(string(parts[1]))
	if err != nil {
		return err
	}

	compareHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if string(compareHash) != string(hash) {
		return fmt.Errorf("password does not match hash")
	}

	return nil
}
