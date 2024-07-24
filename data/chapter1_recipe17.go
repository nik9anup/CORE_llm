// This program demonstrates the use of functional options to configure a Client struct,
// allowing configuration from a JSON file and environmental variables, and prints the resulting configuration.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	consulIP   string
	connString string
}

func (c *Client) String() string {
	return fmt.Sprintf("ConsulIP: %s , Connection String: %s",
		c.consulIP, c.connString)
}

var defaultClient = Client{
	consulIP:   "localhost:9000",
	connString: "postgres://localhost:5432",
}

// ConfigFunc works as a type to be used
// in functional options
type ConfigFunc func(opt *Client)

// FromFile func returns the ConfigFunc
// type. So this way it could read the configuration
// from the json.
func FromFile(path string) ConfigFunc {
	return func(opt *Client) {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		decoder := json.NewDecoder(f)

		fop := struct {
			ConsulIP   string `json:"consul_ip"`
			ConnString string `json:"conn_string"`
		}{}
		err = decoder.Decode(&fop)
		if err != nil {
			panic(err)
		}
		opt.consulIP = fop.ConsulIP
		opt.connString = fop.ConnString
	}
}

// FromEnv reads the configuration
// from the environmental variables
// and combines them with existing ones.
func FromEnv() ConfigFunc {
	return func(opt *Client) {
		connStr, exist := os.LookupEnv("CONN_DB")
		if exist {
			opt.connString = connStr
		}
	}
}

func NewClient(opts ...ConfigFunc) *Client {
	client := defaultClient
	for _, val := range opts {
		val(&client)
	}
	return &client
}

func main() {
	client := NewClient(FromFile("config.json"), FromEnv())
	fmt.Println(client.String())
}