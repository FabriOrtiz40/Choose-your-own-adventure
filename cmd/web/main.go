package main

import (
	"chooseyouradventure/story"
	//"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the story")

	flag.Parse()

	fmt.Printf("Using the story in %s .\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s, err := story.JsonStory(f)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", story.NewHandler(s))
}
