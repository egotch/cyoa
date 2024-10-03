package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/egotch/cyoa"
)

func main() {

  //flags
	filePtr := flag.String("story", "gopher.json", "path + file name of json file containing story data")
  portPtr := flag.Int("port", 3000, "the port to start the CYOA web app on")

	flag.Parse()

	fmt.Printf("\nUsing story file - %s\n", *filePtr)
	file, err := os.Open(*filePtr)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(file)

  h := cyoa.NewHandler(story)
  fmt.Printf("Starting the server on port: %d\n", *portPtr)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), h))

}
