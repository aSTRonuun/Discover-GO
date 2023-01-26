package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Genarate returns a new app ready to be run
func Genarate() *cli.App {

	app := cli.NewApp()

	app.Name = "Command Line App"
	app.Usage = "Find IPs and Servers Names in a website"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Finds the IP of a website",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "server",
			Usage:  "Finds the name servers of a website",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	// net package
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")

	// net package
	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
