package main

//type Messages struct {

//	id, key, value, timestamp string
//}
	

import (
	// "../config.yaml"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type myData struct {
    NumServers int64
    NumFailures int64
    // servers struct [
	// 	id int64
	// 	port int64
	// ]
}


func readConf(filename string) (*myData, error) {
	
    buf, err := ioutil.ReadFile(filename)
	
    if err != nil {
        return nil, err
    }


    c := &myData{}

    err = yaml.Unmarshal(buf, c)
    if err != nil {
        return nil, fmt.Errorf("in file %q: %v", filename, err)
    }

    return c, nil
}

func read() {

}

func write() {

}

func main() {

	// Read YAML file
	c, err := readConf("config.yaml")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%#v", c)

	// extract information

	// Using configuration data, connect all the servers to each other

}
