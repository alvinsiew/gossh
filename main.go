package main

import (
	"flag"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	addAction := flag.String("add", "", "Adding new Hosts")
	listAction := flag.Bool("l", false, "List all hosts")

	flag.Parse()

	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	

	fmt.Println("host:", *addAction)

	listBool := *listAction

	fmt.Println("hello: ",listBool)
	 
	if listBool == true {
		listBucket(db, "GOSSH")
	}
	// abc := flag.Args()
	// fmt.Println(len(abc))
	
	// fmt.Println(abc[1])



	// if lenArgument == 0 {
	// 	listBucket(db, "GOSSH")
	// }

	// fmt.Println(lenArgument)
	// if lenArgument > 0 {
	// 	var firstArg string
	// 	firstArg = os.Args[1]
	// 	if firstArg == "add" {
	// 		if lenArgument != 3 {
	// 			fmt.Println("For add parameter, host and ip address are require only.")
	// 			os.Exit(1)
	// 		}
	// 	}
	// }
	// var secondArg string
	// secondArg = os.Args[2]

	// err = addHost(db, firstArg, secondArg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = addHost(db, "gms-test", "192.168.2.2")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = addHost(db, "testabc", "192.168.3.2")
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
	// 	b := tx.Bucket([]byte("GOSSH"))
	// 	b.ForEach(func(k, v []byte) error {
	// 		fmt.Println(string(k), string(v))
	// 		return nil
	// 	})
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.View(func(tx *bolt.Tx) error {

	// 	c := tx.Bucket([]byte("GOSSH"))
	// 	findHost := c.Get([]byte("testabc"))
	// 	fmt.Printf("Ip address is %s\n", findHost)

	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
		_, err = root.CreateBucketIfNotExists([]byte("HOSTS"))
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

// Function for adding new record
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

// Function for listings all hosts
func listBucket(db *bolt.DB, b string) {

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b))
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
