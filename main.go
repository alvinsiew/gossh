package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	bolt "go.etcd.io/bbolt"
)

// Config struct which contain hosts infomation
type Config struct {
	IP         string `json:"ip"`
	User       string `json:"user"`
	PortNumber string `json:"port"`
	Key        string `json:"key"`
}

func main() {
	addParam := flag.Bool("add", false, "Adding new Hosts")
	ipParam := flag.String("ip", "", "Adding or changing IP address for host")
	userParam := flag.String("user", "", "User")
	portParam := flag.String("port", "22", "Update Port Number. Default(22)")
	keyParam := flag.String("key", "", "Setup key to for server connection. Using default key if not specific.")
	listParam := flag.Bool("l", false, "List all hosts config")

	flag.Parse()

	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// fmt.Println("host:", *addParam, *userParam, *portParam, *ipParam, *keyParam)

	listBool := *listParam
	addBool := *addParam

	if listBool == true {
		listBucket(db, "GOSSH", "CONFIG")
	} else if addBool == true {
		hostArray := flag.Args()
		if len(hostArray) <=0 {
			fmt.Println("Hostname is require")
			os.Exit(1)
		}
		host := hostArray[0]
		fmt.Println(host)
		//

		err = addConfig(db, *ipParam, *userParam, *portParam, *keyParam, host)
		if err != nil {
			// log.Fatal(err)
			fmt.Printf("Error: %v", err)
		}
	}
}

// Function for setting up DB and Bucket
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
		_, err = root.CreateBucketIfNotExists([]byte("CONFIG"))
		if err != nil {
			return fmt.Errorf("could not create weight bucket: %v", err)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	// fmt.Println("DB Setup Done")
	return db, nil
}

// Function for listings all hosts
func listBucket(db *bolt.DB, b string, c string) {

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

// Function for updating database
func addConfig(db *bolt.DB, ip string, user string, port string, key string, hostname string) error {
	config := Config{IP: ip, User: user, PortNumber: port, Key: key}
	configBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("could not marshal config json: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("GOSSH")).Bucket([]byte("CONFIG")).Put([]byte(hostname), configBytes)
		if err != nil {
			return fmt.Errorf("could not insert config: %v", err)
		}

		return nil
	})
	fmt.Printf("Added Config for %v\n", hostname)
	return err
}
