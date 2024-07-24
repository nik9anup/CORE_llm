// This program demonstrates parsing of date and time strings using the time package in Go,
// handling different time zone scenarios with Parse and ParseInLocation functions.

package main

import (
	"fmt"
	"time"
)

func main() {

	// If timezone is not defined, Parse function returns the time in UTC timezone.
	t, err := time.Parse("2/1/2006", "31/7/2015")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// If timezone is given, it is parsed in the specified timezone.
	t, err = time.Parse("2/1/2006  3:04 PM MST", "31/7/2015  1:25 AM DST")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// ParseInLocation parses the time in the given location if the string does not contain time zone definition.
	t, err = time.ParseInLocation("2/1/2006  3:04 PM ", "31/7/2015  1:25 AM ", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

}