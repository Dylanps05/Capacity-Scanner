package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Dylanps05/Capacity-Scanner/internal"
	"github.com/joho/godotenv"
)

func wait_for_quit() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-done
}

func main() {
	godotenv.Load("/usr/local/etc/capacityscanner.env")

	site_address := os.Getenv("LISTEN_ADDRESS")
	db_address := os.Getenv("DB_ADDRESS")
	site := internal.Site{}
	site.Init(site_address, db_address)

	wait_for_quit()
}
