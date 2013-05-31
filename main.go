package main;

import (
  "fmt"
  "flag"
  "os"
  "net/http"
  "io/ioutil"
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

  response, err_read := ioutil.ReadAll(resp.Body)

  if err_read != nil {
    fmt.Printf("Failed to read error body. Reason: \n")
    fmt.Printf("%v \n", err_read)
    os.Exit(3)
  }

  err_io := ioutil.WriteFile(output_file, response, os.FileMode(0777))

  if err_io != nil {
    fmt.Printf("Failed to write file. Reason: \n")
    fmt.Printf("%v \n", err_io)
    os.Exit(4)
  }

  fmt.Printf("\n \n Download Finished! \n")
}
