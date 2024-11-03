package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/service"
	"github.com/robfig/cron"
)

func listen() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	log.Println("Ending cron")
}

func main() {
	c := cron.New()

	log.Println("Starting cron")

	deleteInactiveUsersService := service.NewDeleteInactiveUsersService()

	c.AddFunc("@daily", func() {
		log.Println("Deleting inactive users")
		err := deleteInactiveUsersService.DeleteInactiveUsers()

		if err != nil {
			log.Println(err)
		}

		log.Println("Deleted inactive users")
	})

	go c.Start()

	listen()
}
