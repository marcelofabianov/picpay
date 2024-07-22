package hasher

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	saltLength = 16
	time       = 1
	memory     = 64 * 1024
	threads    = 4
	keyLength  = 32
)

type Config struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

var defaultConfig = Config{
	Time:    time,
	Memory:  memory,
	Threads: threads,
	KeyLen:  keyLength,
}

type Hasher struct {
	Config Config
}

func NewHasher() *Hasher {
	config := Config{
		Time:    time,
		Memory:  memory,
		Threads: threads,
		KeyLen:  keyLength,
	}

	return &Hasher{Config: config}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func (h *Hasher) Hash(data string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(data), salt, defaultConfig.Time, defaultConfig.Memory, defaultConfig.Threads, defaultConfig.KeyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s$%s$%s", "argon2id", b64Salt, b64Hash), nil
}

func (h *Hasher) Compare(data, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 3 {
		return false, fmt.Errorf("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[2])
	if err != nil {
		return false, err
	}

	newHash := argon2.IDKey([]byte(data), salt, defaultConfig.Time, defaultConfig.Memory, defaultConfig.Threads, uint32(len(hash)))

	if subtle.ConstantTimeCompare(hash, newHash) == 1 {
		return true, nil
	}

	return false, nil
}
