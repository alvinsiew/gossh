package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alvinsiew/gossh/internal/gossh"
)

func main() {

	rootBucket := "GOSSH"
	bucket := "HOSTS"
	gosshDB := "gossh.db"

	addParam := flag.Bool("add", false, "Flag for adding new hosts")
	hostParam := flag.String("host", "", "Hostname")
	ipParam := flag.String("ip", "", "Adding or changing IP address for host")
	userParam := flag.String("user", "", "User")
	portParam := flag.String("port", "22", "Update Port Number. Default(22)")
	keyParam := flag.String("key", "", "Setup key to for server connection. Using default key if not specific.")
	listParam := flag.Bool("l", false, "List all hosts config")

	flag.Parse()

	db, err := gossh.SetupDB(gosshDB, rootBucket, bucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// fmt.Println("host:", *addParam, *userParam, *portParam, *ipParam, *keyParam)

	listBool := *listParam
	addBool := *addParam

	if listBool == true {
		gossh.ListBucket(db, rootBucket, bucket)
	} else if addBool == true {
		if *hostParam == "" {
			fmt.Println("Hostname is require")
			os.Exit(1)
		}
		err = gossh.AddHosts(db, rootBucket, bucket, *hostParam, *ipParam, *userParam, *portParam, *keyParam)
		if err != nil {
			// log.Fatal(err)
			fmt.Printf("Error: %v\n", err)
		}
	}
}
