// This program demonstrates how to serialize and deserialize time.Time values using JSON encoding.

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	// Load the Europe/Vienna time zone location
	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	// Create a time instance in the Europe/Vienna time zone
	t := time.Date(2017, 11, 20, 11, 20, 10, 0, eur)

	// Serialize as RFC 3339
	b, err := t.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("Serialized as RFC 3339:", string(b))

	// Deserialize from RFC 3339
	t2 := time.Time{}
	t2.UnmarshalJSON(b)
	fmt.Println("Deserialized from RFC 3339:", t2)

	// Serialize as epoch
	epoch := t.Unix()
	fmt.Println("Serialized as Epoch:", epoch)

	// Deserialize from epoch
	jsonStr := fmt.Sprintf("{ \"created\":%d }", epoch)
	data := struct {
		Created int64 `json:"created"`
	}{}
	json.Unmarshal([]byte(jsonStr), &data)
	deserialized := time.Unix(data.Created, 0)
	fmt.Println("Deserialized from Epoch:", deserialized)
}