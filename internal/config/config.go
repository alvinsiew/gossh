package config

import (
	"os"
	"os/user"
)

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
