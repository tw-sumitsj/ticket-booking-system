package main

import (
  "flag"
  "fmt"
)
func main() {
  migrate := flag.Bool("migrate", false, "To run migrate")
  flag.Parse()

  if *migrate {
    doMigrate()
    return
  }

  doStartServer()
}

func doStartServer() {
  fmt.Println("Start Server..")
}

func doMigrate() {
 fmt.Println("Run Migration")
}
