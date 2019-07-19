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

func TestSetupDBAddHosts(t *testing.T) {
	db, err := SetupDB(dbFile, rootBucketTest, bucketTest)
	if err != nil {
		t.Errorf("Cannot open db")
	}
	defer db.Close()

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		t.Errorf("db file not found")
	}

	findKey := GetKey()
	if len(findKey) <= 0 {
		err := KeyGen()
		if err != nil {
			t.Errorf("Error generating sha key: %v", err)
		}
	}
	
	err = AddHosts(db, rootBucketTest, bucketTest, hostname, ip, user, port, pass, key)
	if err != nil {
		t.Errorf("Error inserting value %v", err)
	}

	h := FindHost(db, rootBucketTest, bucketTest, hostname)
	if h.IP != ip {
		t.Errorf("Ip address does not match %s: %s", ip, h.IP)
	}
}
