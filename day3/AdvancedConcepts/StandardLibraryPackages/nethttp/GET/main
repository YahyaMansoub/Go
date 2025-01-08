package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Define the URL you want to send a GET request to
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Send the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err) // Log and exit if there is an error with the request
	}
	defer resp.Body.Close() // Ensure that the body is closed when done

	// Check if the response status is OK (200)
	if resp.StatusCode == http.StatusOK {
		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Print the response body as a string
		fmt.Println("Response Body:")
		fmt.Println(string(body))
	} else {
		// If status code is not 200, print an error message
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	}
}
