package main

import (
  "fmt"
  "os"
  "log"
  "github.com/joho/godotenv"
)

func main() {
  fmt.Println("hello world")

  // by default it extract from file .env but we can specified the name file as well -> godotenv.Load(".env")
  godotenv.Load()  

  portString := os.Getenv("PORT")  // to get value of var PORT in the current shell env
  if portString == "" {
    log.Fatal("PORT is not found in env")
  }

  fmt.Println("Port:", portString)
}