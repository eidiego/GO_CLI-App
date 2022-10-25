package app

import "github.com/urfave/cli"

// Generate will return the CLI application ready to execute
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI application"
	app.Usage = "Search IPs and server names in the internet"

	return app
}
