package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "net/http"
//"os"
//"strings"
  "regexp"
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

  //Read Response data into memory
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal("Error reading HTTP body. ", err)
  }

  //Create Regex to find the captions
  re := regexp.MustCompile(`(?i)<p class="">(.+)</p>`)
  captions := re.FindAllStringSubmatch(string(body), -1)
  if captions == nil {
    fmt.Println("No matches.")
  } else {
      for _, caption := range captions {
        fmt.Println(caption[1])
      }
  }

}
