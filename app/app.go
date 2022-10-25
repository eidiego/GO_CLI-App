package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the CLI application ready to execute
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI application"
	app.Usage = "Search IPs and server names in the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search address IPs in the internet",
			Flags:  flags,
			Action: searchIps,
		},

		{
			Name:   "servers",
			Usage:  "Search the name servers of the domain in the internet",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
