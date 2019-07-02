package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

// Entry type
type Entry struct {
	Host      string `json:"host"`
	IPaddress string `json:"ip"`
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = addHost(db, "host_abc", "192.168.1.2")
	if err != nil {
		log.Fatal(err)
	}

	err = addHost(db, "gms-test", "192.168.2.2")
	if err != nil {
		log.Fatal(err)
	}

	err = addHost(db, "testabc", "192.168.3.2")
	if err != nil {
		log.Fatal(err)
	}

	// err = addEntry(db, "testabc", "192.168.2.3", time.Now())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Manipulating the date
	// err = addEntry(db, "testcde", "172.134.2.32", time.Now().AddDate(0, 0, -2))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("GOSSH")).Bucket([]byte("HOSTS"))
	// 	b.ForEach(func(k, v []byte) error {
	// 		fmt.Println(string(k), string(v))
	// 		return nil
	// 	})
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = db.View(func(tx *bolt.Tx) error {

		c := tx.Bucket([]byte("GOSSH"))
		// find := c.Seek([]byte
		findHost := c.Get([]byte("testabc"))
		fmt.Printf("Ip address is %s\n", findHost)
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

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("gossh.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("GOSSH"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("HOSTS"))
		if err != nil {
			return fmt.Errorf("could not create weight bucket: %v", err)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

func addHost(db *bolt.DB, host string, ip string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("GOSSH")).Put([]byte(host), []byte(ip))
		if err != nil {
			return fmt.Errorf("could not insert hosts table: %v", err)
		}
		return nil
	})
	fmt.Println("Added host")
	return err
}

func addEntry(db *bolt.DB, host string, ip string, date time.Time) error {
	entry := Entry{Host: host, IPaddress: ip}
	entryBytes, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("could not marshal entry json: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("GOSSH")).Bucket([]byte("HOSTS")).Put([]byte(date.Format(time.RFC3339)), entryBytes)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}

		return nil
	})
	fmt.Println("Added Entry")
	return err
}
