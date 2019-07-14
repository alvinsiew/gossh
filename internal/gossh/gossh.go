package gossh

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/alvinsiew/gossh/pkg/sshclient"
	"github.com/alvinsiew/gossh/internal/crypto"
	bolt "go.etcd.io/bbolt"
)

// Config struct which contain hosts infomation
type Config struct {
	IP         string `json:"ip"`
	User       string `json:"user"`
	PortNumber string `json:"port"`
	Password   string `json:"password"`
	Key        []byte `json:"key"`
}

// SetupDB for setting up root bucket and bucket
func SetupDB(dbFile string, rootBucket string, bucket string) (*bolt.DB, error) {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte(rootBucket))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("could not create %s bucket: %v", bucket, err)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	return db, nil
}

// ListBucket for listings all hosts and value
func ListBucket(db *bolt.DB, rootBucket string, bucket string) {
	err := db.View(func(tx *bolt.Tx) error {
		rootBucket := tx.Bucket([]byte(rootBucket)).Bucket([]byte(bucket))
		rootBucket.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// ListBucketHOSTS for listings all hosts
func ListBucketHOSTS(db *bolt.DB, rootBucket string, bucket string) {
	err := db.View(func(tx *bolt.Tx) error {
		rootBucket := tx.Bucket([]byte(rootBucket)).Bucket([]byte(bucket))
		rootBucket.ForEach(func(k, v []byte) error {
			fmt.Println(string(k))
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// AddHosts for updating database with hosts informations
func AddHosts(db *bolt.DB, rootBucket string, bucket string, hostname string, ip string, user string, port string, pass string, key []byte) error {
	config := Config{IP: ip, User: user, PortNumber: port, Password: pass, Key: key}
	configBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("could not marshal config json: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(rootBucket)).Bucket([]byte(bucket)).Put([]byte(hostname), configBytes)
		if err != nil {
			return fmt.Errorf("could not insert %v: %v", bucket, err)
		}

		return nil
	})
	fmt.Printf("Updated %v for %v\n", bucket, hostname)
	return err
}

// AddConf for updating database with hosts informations
func AddConf(db *bolt.DB, rootBucketConf string, key string, value string) error {
	shaKey := crypto.CreateHash(key)
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(rootBucketConf)).Put([]byte(key), []byte(shaKey))
		if err != nil {
			return fmt.Errorf("could not insert %v", err)
		}

		return nil
	})
	return err
}

// FindHost search for host detail
func FindHost(db *bolt.DB, rootBucket string, bucket string, host string) Config {
	var c Config
	err := db.View(func(tx *bolt.Tx) error {
		hostDetails := tx.Bucket([]byte(rootBucket)).Bucket([]byte(bucket)).Get([]byte(host))
		if hostDetails == nil {
			fmt.Printf("Unable to find host %s\n", host)
			os.Exit(1)
		}
		err := json.Unmarshal([]byte(hostDetails), &c)
		if err != nil {
			log.Fatal(err)
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// FindConf search for key pair
func FindConf(db *bolt.DB, rootBucket string, key string) string {
	var result []byte
	err := db.View(func(tx *bolt.Tx) error {
		hostDetails := tx.Bucket([]byte(rootBucket))
		result = hostDetails.Get([]byte(key))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

// DeleteHost for deleting host from db
func DeleteHost(db *bolt.DB, rootBucket string, bucket string, host string) error {

	err := db.Update(func(tx *bolt.Tx) error {
		hostDetails := tx.Bucket([]byte(rootBucket)).Bucket([]byte(bucket))
		err := hostDetails.Delete([]byte(host))
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// SSSHConn make ssh connection to server
func (c Config) SSSHConn() {
	ip := c.IP
	user := c.User
	port := c.PortNumber
	key := c.Key
	pass := c.Password

	err := sshclient.TerminalConn(user, key, ip, port, pass)
	if err != nil {
		log.Fatalf("Fail connection %s", err)
	}
}
