package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bertinatto/leapp-ng/pkg/web"
)

var (
	flagVersion = flag.Bool("version", false, "show version")
	flagHelp    = flag.Bool("help", false, "show usage")
)

func init() {
	//if debug, _ := strconv.ParseBool(os.Getenv("CAMLI_DEBUG")); debug {

	//}
}

func main() {
	// this way we can use Main() in tests
	os.Exit(Main())
}

func Main() int {
	webHandler := web.New()
	go webHandler.Run()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Println("Received SIGTERM. Shutting down gracefully...")
	case err := <-webHandler.ErrorCh():
		log.Printf("Error starting service: %v\n", err)
		return 1

	}

	return 0

}
