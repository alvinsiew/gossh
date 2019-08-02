package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alvinsiew/gossh/internal/config"
	"github.com/alvinsiew/gossh/internal/gossh"
	"github.com/alvinsiew/gossh/pkg/sshclient"
)

var rootBucket = "GOSSH"
var bucket = "HOSTS"
var gosshDB = "gossh.db"
var gosshCONF = "conf.db"
var defaultUser = *config.GetCurrentUser()
var defaultHome = defaultUser.HomeDir
var gosshDir = defaultHome + "/.gossh/"
var dbConfPath = gosshDir + gosshCONF
var gosshDBpath = gosshDir + gosshDB

func init() {
	defaultUser := *config.GetCurrentUser()
	defaultHome := defaultUser.HomeDir
	gosshDir := defaultHome + "/.gossh/"
	config.MakeDir(gosshDir)

	findKey := gossh.GetKey(dbConfPath)
	if len(findKey) <= 0 {
		err := gossh.KeyGen(dbConfPath)
		if err != nil {
			log.Fatalf("Error generating sha key: %s", err)
		}
	}
}

func main() {
	addParam := flag.Bool("add", false, "Add host:\nUsage: gossh -add -host <hostname|mandatory> -ip <ip address|mandatory> -user <userid|non-mandatory> -port <ssh port|non-mandatory> -key <private key|non-mandatory>")
	delParam := flag.Bool("del", false, "Hostname to delete")
	hostParam := flag.String("host", "", "Hostname")
	ipParam := flag.String("ip", "", "Adding or changing IP address for host")
	userParam := flag.String("user", defaultUser.Username, "User")
	portParam := flag.String("port", "22", "Port Number")
	passParam := flag.String("pass", "", "User password")
	keyParam := flag.String("key", "nokey", "Setup key to for server connection. Using default key if not specific.")
	listParam := flag.Bool("l", false, "List all hosts config\n -l info \nto list more infor\n -l key \nto list private key")
	connParam := flag.Bool("c", false, "Connection to server:\nUsage: gossh -c <hostname>\n")

	flag.Parse()

	db, err := gossh.SetupDB(gosshDBpath, rootBucket, bucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if *listParam == true {
		n := len(flag.Args())
		if n == 0 {
			gossh.ListBucketHOSTS(db, rootBucket, bucket)
		} else if flag.Args()[0] == "info" {
			gossh.ListBucket(db, rootBucket, bucket, dbConfPath, flag.Args()[0])
		} else if flag.Args()[0] == "key" {
			gossh.ListBucket(db, rootBucket, bucket, dbConfPath, flag.Args()[0])
		} else {
			log.Fatalf("Invalid argument")
		}
	} else if *addParam == true {
		if *hostParam == "" {
			fmt.Println("Hostname is require")
			os.Exit(1)
		} else if *keyParam != "nokey" {
			keyBytes := sshclient.KeyToBytes(*keyParam)
			err = gossh.AddHosts(db, rootBucket, bucket, dbConfPath, *hostParam, *ipParam, *userParam, *portParam, *passParam, keyBytes)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else if *keyParam == "nokey" {
			keyBytes := make([]byte, 0)
			err = gossh.AddHosts(db, rootBucket, bucket, dbConfPath, *hostParam, *ipParam, *userParam, *portParam, *passParam, keyBytes)
		}
	} else if *connParam == true {
		n := len(flag.Args())
		if n == 0 {
			fmt.Printf("No hostname to connect\n")
			os.Exit(1)
		}
		host := flag.Args()[0]
		result := gossh.FindHost(db, rootBucket, bucket, dbConfPath, host)
		result.SSHConn()
	} else if *delParam == true {
		n := len(flag.Args())
		if n == 0 {
			fmt.Printf("No hostname to delete\n")
			os.Exit(1)
		}
		host := flag.Args()[0]
		err := gossh.DeleteHost(db, rootBucket, bucket, host)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}
}
