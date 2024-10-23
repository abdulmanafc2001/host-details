package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		return
	}

	// Get the list of network interfaces
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Printf("Error looking up hostname: %v\n", err)
		return
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("HostName: %s, IpAddr: %v\n", hostname, addrs))
	})

	app.Listen(":3001")
}
