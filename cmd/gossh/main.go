package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alvinsiew/gossh/internal/gossh"
	"github.com/alvinsiew/gossh/pkg/sshclient"
)

func main() {
	rootBucket := "GOSSH"
	bucket := "HOSTS"
	gosshDB := "gossh.db"

	addParam := flag.Bool("add", false, "Add host:\nUsage: gossh -host <hostname|mandatory> -ip <ip address|mandatory> -user <userid|non-mandatory> -port <ssh port|non-mandatory> -key <private key|non-mandatory>")
	hostParam := flag.String("host", "", "Hostname")
	ipParam := flag.String("ip", "", "Adding or changing IP address for host")
	userParam := flag.String("user", "", "User")
	portParam := flag.String("port", "22", "Update Port Number. Default(22)")
	keyParam := flag.String("key", "nokey", "Setup key to for server connection. Using default key if not specific.")
	listParam := flag.Bool("l", false, "List all hosts config")
	connParam := flag.Bool("conn", false, "Connection to server:\nUsage: gossh -conn <hostname>\n")

	flag.Parse()

	db, err := gossh.SetupDB(gosshDB, rootBucket, bucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println(*keyParam)
	if *listParam == true {
		gossh.ListBucket(db, rootBucket, bucket)
	} else if *addParam == true {
		if *hostParam == "" {
			fmt.Println("Hostname is require")
			os.Exit(1)
		} else if *keyParam != "nokey" {
			fmt.Println(*keyParam)
			keyBytes := sshclient.KeyToBytes(*keyParam)
			err = gossh.AddHosts(db, rootBucket, bucket, *hostParam, *ipParam, *userParam, *portParam, keyBytes)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else if *keyParam == "nokey" {
			keyBytes := make([]byte, 0)
			// fmt.Println(keyBytes)
			err = gossh.AddHosts(db, rootBucket, bucket, *hostParam, *ipParam, *userParam, *portParam, keyBytes)
		}
	} else if *connParam == true {
		n := len(flag.Args())
		if n == 0 {
			fmt.Printf("No hostname to connect\n")
			os.Exit(1)
		}
		host := flag.Args()[0]
		result := gossh.FindHost(db, rootBucket, bucket, host)
		result.SSSHConn()
	}
}