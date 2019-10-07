package main

import (
	"flag"
	"fmt"
	"goplay/cyoa"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	storyFileName := flag.String("storyFileName", "gopher.json", "the json file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *storyFileName)

	file, err := os.Open(*storyFileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story) //Verbose print statement

	// Create our custom CYOA story handler
	h := cyoa.NewHandler(story)
	// Create a ServeMux to route our requests
	mux := http.NewServeMux()
	// This story handler is using a custom function and template
	// Because we use /story/ (trailing slash) all web requests
	// whose path has the /story/ prefix will be routed here.
	mux.Handle("/story/", h)
	// This story handler is using the default functions and templates
	// Because we use / (base path) all incoming requests not
	// mapped elsewhere will be sent here.
	mux.Handle("/", cyoa.NewHandler(story))
	// Start the server using our ServeMux
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}