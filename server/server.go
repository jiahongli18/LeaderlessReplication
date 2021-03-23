package main

import (
	"LeaderlessReplication/receiver"
	"LeaderlessReplication/sender"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	NumServers  int
	NumFailures int
	Servers     []Server
}

type Server struct {
	ID   int
	IP   string
	Port string
}

type Data struct {
	Key   string
	Value string
}

//global data storage
var data []Data

// readConf reads in YAML file and creates Config data
func readConf(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	c := &Config{}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func main() {

	// takes in the command argument and identify the ID and the port number
	arguments := os.Args

	// checks if host address and port # are provided
	if len(arguments) == 1 {
		fmt.Println("Please provide port")
		return
	}

	port := arguments[1]

	// reads YAML file and extract information
	c, err := readConf("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// get ID of this server from config file and port provided
	var ID int

	for i := 0; i < c.NumServers; i++ {
		// fmt.Println(c.Servers[i])
		if c.Servers[i].Port == port {
			ID = c.Servers[i].ID
		}
	}

	nodes := make(map[string]net.Conn) // tracks server ID and InetAddress

	// listens
	for i := ID + 1; i < c.NumServers; i++ {
		fmt.Println("listening for server: ", ID, "port: ", port)
		receiver.ServerListen(port, nodes)
	}


	for i := 0; i < ID; i++ {
			fmt.Println("dialing: ", i)
			// sender.Dial(i, ID, IDs, IPs, ports, nodes)
			sender.Dial(i, ID, nodes)
	}

	// for {
	//   // loop through c.Servers and c
	//   // N - f
	// }
}
