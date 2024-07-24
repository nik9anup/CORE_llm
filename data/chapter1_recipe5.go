// This code looks up an environment variable and logs a message if it is not set, then prints the value of the environment variable.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	connStr, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("The env variable %s is not set.\n", key)
	}
	fmt.Println(connStr)
}