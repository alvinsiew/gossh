package gossh

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

// FindHost search for host detail
func FindHost(db *bolt.DB, rootBucket string, bucket string, host string) {

		// fmt.Printf("Test %v\n", rootBucket)
	err := db.View(func(tx *bolt.Tx) error {
		hostDetails := tx.Bucket([]byte(rootBucket)).Bucket([]byte(bucket)).Get([]byte(host))

		// c := tx.Bucket([]byte(rootBucket))
		// fmt.Println(&c)
		// find := c.Seek([]byte
		// findHost := c.Get([]byte(host))
		// fmt.Println(findHost)
		fmt.Printf("%s\n", hostDetails)
		// min := []byte(time.Now().AddDate(0, 0, -7).Format(time.RFC3339))
		// max := []byte(time.Now().AddDate(0, 0, 0).Format(time.RFC3339))
		// for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
		// 	fmt.Println(string(k), string(v))
		// }
		// c.Bucket([]byte("HOSTS"))).Get
		// findHost := c.
		// fmt.Println(c.Seek(findHost))

		// for k, v := c.Seek()
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}