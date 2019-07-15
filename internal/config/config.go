package config

import (
	"math/rand"
	"os"
	"os/user"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetCurrentUser for checking current user
func GetCurrentUser() *user.User {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	return user
}

// MakeDir for making directory
func MakeDir(d string) {
	if _, err := os.Stat(d); os.IsNotExist(err) {
		os.Mkdir(d, 0700)
	}
}

// RandStringRunes generate random word
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
