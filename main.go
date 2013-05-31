package main;

import (
  "fmt"
  "flag"
  "os"
)

func main() {
  var address string
  var number_of_connections int

  flag.StringVar(&address, "a", "", "The address to download the file")
  flag.IntVar(&number_of_connections, "n", 1, "Number of connections to download the file")

  flag.Parse()

  fmt.Printf("YADM: Yet another download manager \n")
  fmt.Printf("By Thiago Pradi (thiago.pradi@gmail.com) \n \n")

  if address == "" {
    fmt.Printf("You need a url to start downloading! \n")
    os.Exit(1)
  }

  fmt.Printf("Address: %v \n", address)
  fmt.Printf("Number of connections: %d \n \n", number_of_connections)
}
