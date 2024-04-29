package shorter

import (
	config "github.com/roman127001/urlsha"
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Shorter interface {
	Generate() string
}

func New() *Client {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Client{}
}

type Client struct{}

func (s *Client) Generate() string {
	key := make([]byte, config.ShortUrlLength)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}

	return string(key)
}
