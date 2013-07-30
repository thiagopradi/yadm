package main;

import (
  "fmt"
  "flag"
  "os"
  "net/http"
  "io"
)

func makeRequest(address string, start_byte int64, end_byte int64, out *os.File) {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", address, nil)
  header_string := fmt.Sprintf("bytes=%d-%d", start_byte, end_byte)
  fmt.Printf("Header String %v \n", header_string)

  req.Header.Set("Range", header_string)
  res, err := client.Do(req)

  if err != nil {
    fmt.Printf("Download failed. Reason: \n")
    fmt.Printf("%v \n", err)
    os.Exit(2)
  } else {
    defer res.Body.Close()
    io.Copy(out, res.Body)
  }
}

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

  head_resp, head_err := http.Head(address)

  if head_err != nil {
    fmt.Printf("Download failed. Reason: \n")
    fmt.Printf("%v \n", head_err)
    os.Exit(2)
  }

  fmt.Printf("File Size: %d Bytes \n", head_resp.ContentLength)

  out, err_create := os.Create(output_file)

  if err_create != nil {
    fmt.Printf("Failed to create file. Reason: \n")
    fmt.Printf("%v \n", err_create)
    os.Exit(3)
  }

  start_byte := int64(0)
  section_number := 0
  section_size := int64(head_resp.ContentLength) / int64(number_of_connections)
  /* channels := make(chan *http.Response, number_of_connections) */
  
  for(section_number < number_of_connections) {
    end_byte := start_byte + section_size
    section_number = section_number + 1

    if int64(head_resp.ContentLength) - end_byte < section_size { 
      end_byte = int64(head_resp.ContentLength)
    }

    fmt.Printf("Section %d: From %d to %d Bytes \n", section_number, start_byte, end_byte)

    makeRequest(address, start_byte, end_byte, out)

    start_byte = start_byte + section_size + 1
  }

  fmt.Printf("\n \n Download Finished! \n")
}
