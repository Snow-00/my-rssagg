package main

import (
  "os"
  "log"
  "net/http"
  "github.com/joho/godotenv"
  "github.com/go-chi/chi"
)

func main() {
  // by default it extract from file .env but we can specified the name file as well -> godotenv.Load(".env")
  godotenv.Load()  

  portString := os.Getenv("PORT")  // to get value of var PORT in the current shell env
  if portString == "" {
    log.Fatal("PORT is not found in env")
  }

  router := chi.NewRouter()

  srv := &http.Server{
    Handler: router,
    Addr: ":" + portString,
  }
  
  log.Printf("server starting on port %v", portString)
  err := srv.ListenAndServe()
  if err != nil{
    log.Fatal(err)
  }
}