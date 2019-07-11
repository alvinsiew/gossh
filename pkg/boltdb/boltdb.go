package boltdb

import (
	"encoding/json"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

// Config struct which contain hosts infomation
type Config struct {
	IP         string `json:"ip"`
	User       string `json:"user"`
	PortNumber string `json:"port"`
	Key        string `json:"key"`
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
	// fmt.Println("DB Setup Done")
	return db, nil
}

// ListBucket for listings all hosts
func ListBucket(db *bolt.DB, b string, c string) {

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b)).Bucket([]byte(c))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// AddHosts for updating database with hosts informations
func AddHosts(db *bolt.DB, rootBucket string, bucket string, hostname string, ip string, user string, port string, key string) error {
	config := Config{IP: ip, User: user, PortNumber: port, Key: key}
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
	fmt.Printf("Added %v for %v\n", bucket, hostname)
	return err
}
