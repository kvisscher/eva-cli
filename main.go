package main

import (
	"fmt"
	"os"

	"github.com/new-black/eva-cli/client"
	"github.com/new-black/eva-cli/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	client := client.NewHttpClient("https://eva.newblack.io")

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "auth",
			Subcommands: cmd.GenerateLoginCommands(client),
		},
		cli.Command{
			Name:        "applications",
			Subcommands: cmd.GenerateApplicationCommands(client),
		},
		cli.Command{
			Name:        "blob",
			Subcommands: cmd.GenerateBlobCommands(client),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%+v", err)
		fmt.Println()
	}
}
