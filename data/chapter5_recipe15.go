// Parses JSON data from a constant string js representing an array of User objects,
// using json.NewDecoder to decode each object into a User struct and prints the slice of users.
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const js = `
	[{
		"name":"Axel",
		"lastname":"Fooley"
	},
	{
		"name":"Tim",
		"lastname":"Burton"
	},
	{
		"name":"Tim",
		"lastname":"Burton"
`

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func main() {

	userSlice := make([]User, 0)
	r := strings.NewReader(js)
	dec := json.NewDecoder(r)
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case json.Delim:
			str := tp.String()
			if str == "[" || str == "{" {
				for dec.More() {
					u := User{}
					err := dec.Decode(&u)
					if err == nil {
						userSlice = append(userSlice, u)
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Println(userSlice)
}