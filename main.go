package main

import (
	"github.com/RobinUS2/dns-over-https-balancer/service"
	"log"
)

func main() {
	instance := service.New()
	if err := instance.Run(); err != nil {
		log.Fatalf("failed to start: %s", err)
	}
}
