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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search address IPs in the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: searchIps,
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
