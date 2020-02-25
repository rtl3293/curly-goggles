package main

import (
  "io/ioutil"
  "log"
  "net/http"
//  "os"
  //"strings"
//  "regexp"
  "time"
)

func main() {
  //Create HTTP Client with timeout
  client := &http.Client{
    Timeout: 30 * time.Second,
  }

  //Create and modify HTTP request before being sent
  request, err :=  http.NewRequest("GET", "https://www.rightstorickysanchez.com/normal-column", nil)
  if err != nil {
    log.Fatal(err)
  }
  request.Header.Set("User-Agent", "Lickface")

  //Make the request
  response, err := client.Do(request)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  //Copy data from the response to standard output
  //n, err := io.Copy(os.Stdout, response.Body)
  //if err != nil {
  //  log.Fatal(err)
  //}
  dataInBytes, err := ioutil.ReadAll(response.Body)
  pageContent := string(dataInBytes)


  log.Println(pageContent)
}
