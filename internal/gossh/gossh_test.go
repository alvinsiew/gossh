package gossh

import (
	"os"
	"testing"
)

var rootBucketTest = "root"
var bucketTest = "TEST"
var dbFile = "test.db"
var hostname = "testapp"
var ip = "192.168.2.2"
var user = "centos"
var port = "22"
var pass = "abc123"
var key = []byte{1,2,3,4,5,6,7,8}
var dbConf = "testconf.db"

func TestSetupDBAddHostsGetKeyAddHostsFindHost(t *testing.T) {
	db, err := SetupDB(dbFile, rootBucketTest, bucketTest)
	if err != nil {
		t.Errorf("Cannot open db")
	}
	defer db.Close()

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		t.Errorf("db file not found")
	}

	findKey := GetKey(dbConf)
	if len(findKey) <= 0 {
		err := KeyGen(dbConf)
		if err != nil {
			t.Errorf("Error generating sha key: %v", err)
		}
	}
	
	err = AddHosts(db, rootBucketTest, bucketTest, dbConf, hostname, ip, user, port, pass, key)
	if err != nil {
		t.Errorf("Error inserting value %v", err)
	}

	h := FindHost(db, rootBucketTest, bucketTest, dbConf, hostname)
	if h.IP != ip {
		t.Errorf("Ip address does not match %s: %s", ip, h.IP)
	} else if h.PortNumber != port {
		t.Errorf("Port does not match %s: %s", port, h.PortNumber)
	} else if h.User != user {
		t.Errorf("User does not match %s: %s", user, h.User )
	} else if h.Password != pass {
		t.Errorf("Password does not match %s: %s", pass, h.Password)
	}
	os.Remove(dbFile)
	os.Remove(dbConf)
}
