package config

import (
	"testing"
	"os"
)

func TestGetCurrentUser(t *testing.T) {
	getuser := GetCurrentUser()
	
	if len(getuser.Username) <= 0 {
		t.Errorf("Username should not be blank %v", getuser.Username )
	}
}

func TestRandStringRunes(t *testing.T) {
	s := RandStringRunes(10) 
	if len(s) != 10 {
		t.Errorf("Length of string should be 10, but got %v", len(s))
	}
}

func TestMkDir(t *testing.T) {
	MakeDir("testFolder")
	if _, err := os.Stat("testFolder"); os.IsNotExist(err) {
		t.Errorf("Cannot mkdir testFolder")
	}
	os.Remove("testFolder")
}