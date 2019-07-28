package crypto

import (
	"encoding/json"
	"testing"
)

var p = "helloworld"

type Config struct {
	IP         string `json:"ip"`
	User       string `json:"user"`
	PortNumber string `json:"port"`
	Password   string `json:"password"`
	Key        []byte `json:"key"`
}

func TestCrypto(t *testing.T) {
	var c Config
	sha := CreateHash("abcdefghij")
	testData := Config{IP: "192.168.1.120", User: "centos", PortNumber: "22", Password: p, Key: []byte{1, 2, 3, 4, 5, 6, 7, 8}}
	testBytes, err := json.Marshal(testData)
	if err != nil {
		t.Errorf("could not marshal config json: %v", err)
	}

	encData := Encrypt(testBytes, sha)

	decryData, err := Decrypt(encData, sha)
	if err != nil {
		t.Errorf("Could not decrypt data: %v", err)
	}
	err = json.Unmarshal([]byte(decryData), &c)
	if err != nil {
		t.Errorf("Could not unmarshal test data: %v", err)
	}

	if c.IP != "192.168.1.120" {
		t.Errorf("IP Address 192.168.1.120 does not match: %v", c.IP)
	} else if c.User != "centos" {
		t.Errorf("User centos does not match: %v", c.User)
	} else if c.PortNumber != "22" {
		t.Errorf("Port number 22 does not match: %v", c.PortNumber)
	} else if c.Password != p {
		t.Errorf("Password %v does not match: %v", p, c.Password)
	} else if c.Key[2] != 3 {
		t.Errorf("Key 3 does not match: %v", c.Key[2])
	}
}
