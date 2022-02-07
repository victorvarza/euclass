package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {
	l := log.New(os.Stdout, "TEST", log.Lshortfile|log.Ldate|log.Ltime)

	l.Println("Starting Service")
	l.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	defer l.Println("Stopped Service")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
}
