package main

import (
  "fmt"
  "os"
  "log"
)

func main() {
  fmt.Println("hello world")

  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT is not found in env")
  }

  fmt.Println("Port:", portString)
}