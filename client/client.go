package main

import (
	"fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "log"
)

type Config struct {
    NumServers int
    NumFailures int
    Servers []Server
}

type Server struct {
		Id int
		Port int
}

func main() {
  // dials the appropriate server 

  // infinite loop until user exits
  for {
    // interface:
    // 1. get data from the server
    // 2. write key, value to the server
    printInterface()

    input := 

    if input == 1 {

    }

    if input == 2 {

    } else {

    }
  }
}