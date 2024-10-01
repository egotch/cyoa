package main

import (
	"flag"
	"fmt"
	"os"

  "github.com/egotch/cyoa"
)

func main() {

  filePtr := flag.String("story", "gopher.json", "path + file name of json file containing story data")

  flag.Parse()

  fmt.Printf("\nUsing story file - %s\n", *filePtr)
  file, err := os.Open(*filePtr)
  if err != nil {
    panic(err)
  }

  story, err := cyoa.JsonStory(file)

  fmt.Printf("%+v\n", story)
}
