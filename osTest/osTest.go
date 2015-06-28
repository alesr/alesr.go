package main

import (
        "os/exec"
        "log"
        "fmt"
      )

func main() {
  openCaffeinated()
}

func openCaffeinated()  {

  cmd := exec.Command("cmd", "start","/c/Users/Alessandro/Apps/Caffeinated.exe")
  err := cmd.Start()
  if err != nil {
    log.Fatal(err)
  }
  //err = cmd.Process.Kill()
  coffeeCh := make(chan bool, 1)
  coffeeCh <- true
  closeCaffeinated(coffeeCh)
}

func closeCaffeinated(cCh chan bool) {
  status := <- cCh
  cCh <- false
  status2 := <- cCh

  fmt.Println(status)
  fmt.Println(status2)
}
