package shorter

import (
	config "github.com/roman127001/urlsha"
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func New() *Shorter {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Shorter{}
}

type Shorter struct{}

func (s *Shorter) Generate() string {
	key := make([]byte, config.ShortUrlLength)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}

	return string(key)
}
