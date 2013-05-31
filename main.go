package main;

import (
  "fmt"
  "flag"
  "os"
  "net/http"
  "io"
)

func main() {
  var address string
  var output_file string
  var number_of_connections int

  flag.StringVar(&address, "a", "", "The address to download the file")
  flag.StringVar(&output_file, "o", "output.txt", "The default output file")
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

  resp, err := http.Get(address)

  if err != nil {
    fmt.Printf("Download failed. Reason: \n")
    fmt.Printf("%v \n", err)
    os.Exit(2)
  }

  defer resp.Body.Close()

  out, err_create := os.Create(output_file)

  if err_create != nil {
    fmt.Printf("Failed to create file. Reason: \n")
    fmt.Printf("%v \n", err_create)
    os.Exit(3)
  }

  io.Copy(out, resp.Body)

  fmt.Printf("\n \n Download Finished! \n")
}
