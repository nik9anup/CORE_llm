/*
Package main demonstrates basic CRUD operations (Create, Read) using an HTTP server and client in Go.

This program sets up an HTTP server that manages a list of cities. It allows clients to retrieve the list of cities
via a GET request, and add a new city via a POST request. It also provides a client interface to fetch the list of cities
and save a new city to the server.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// addr represents the address and port on which the HTTP server listens.
const addr = "localhost:7070"

// City represents a city with its ID, name, and location.
type City struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// toJson converts a City struct to its JSON string representation.
func (c City) toJson() string {
	return fmt.Sprintf(`{"name":"%s","location":"%s"}`, c.Name, c.Location)
}

func main() {
	// Create and start the HTTP server in a separate goroutine.
	s := createServer(addr)
	go s.ListenAndServe()

	// Retrieve the list of cities from the server.
	cities, err := getCities()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Retrieved cities: %v\n", cities)

	// Save a new city "Paris" to the server.
	city, err := saveCity(City{"", "Paris", "France"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saved city: %v\n", city)
}

// saveCity sends a POST request to save a new city to the server.
// It returns the saved City object and any error encountered.
func saveCity(city City) (City, error) {
	r, err := http.Post("http://"+addr+"/cities", "application/json", strings.NewReader(city.toJson()))
	if err != nil {
		return City{}, err
	}
	defer r.Body.Close()
	return decodeCity(r.Body)
}

// getCities sends a GET request to retrieve the list of cities from the server.
// It returns a slice of City objects and any error encountered.
func getCities() ([]City, error) {
	r, err := http.Get("http://" + addr + "/cities")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return decodeCities(r.Body)
}

// decodeCity decodes a JSON-encoded City object from the provided reader.
// It returns the decoded City object and any error encountered during decoding.
func decodeCity(r io.Reader) (City, error) {
	city := City{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&city)
	return city, err
}

// decodeCities decodes a JSON-encoded slice of City objects from the provided reader.
// It returns the decoded slice of City objects and any error encountered during decoding.
func decodeCities(r io.Reader) ([]City, error) {
	cities := []City{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&cities)
	return cities, err
}

// createServer creates and returns an HTTP server configured to handle city-related requests.
// It initializes with a predefined set of cities and uses a multiplexer (mux) to route requests.
func createServer(addr string) http.Server {
	// Predefined list of cities.
	cities := []City{
		{ID: "1", Name: "Prague", Location: "Czechia"},
		{ID: "2", Name: "Bratislava", Location: "Slovakia"},
	}

	// Create a new HTTP multiplexer (mux).
	mux := http.NewServeMux()

	// Define handler for "/cities" endpoint.
	mux.HandleFunc("/cities", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		if r.Method == http.MethodGet {
			// Handle GET request: Return the list of cities.
			enc.Encode(cities)
		} else if r.Method == http.MethodPost {
			// Handle POST request: Add a new city to the list.
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			r.Body.Close()

			// Decode the incoming JSON data into a City object.
			city := City{}
			if err := json.Unmarshal(data, &city); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Assign a new ID (incremental) to the new city.
			city.ID = strconv.Itoa(len(cities) + 1)

			// Append the new city to the list of cities.
			cities = append(cities, city)

			// Return the newly added city as a JSON response.
			enc.Encode(city)
		}
	})

	// Create and return an HTTP server configured with the multiplexer (mux).
	return http.Server{
		Addr:    addr,
		Handler: mux,
	}
}