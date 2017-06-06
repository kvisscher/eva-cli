package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"io/ioutil"

	"log"

	"github.com/new-black/eva-cli/client"
	"github.com/new-black/eva-cli/cmd"
	"github.com/urfave/cli"
)

func main() {
	hostFile, err := getPathToHost()

	if err != nil {
		fmt.Printf("%+v", err)
		fmt.Println()
	}

	b, err := ioutil.ReadFile(hostFile)

	if err != nil {
		log.Fatal(err)
	}

	host := string(b)

	if host == "" {
		host = "https://eva.newblack.io"
	}

	evaClient := client.NewHttpClient(host)

	app := cli.NewApp()

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "auth",
			Subcommands: cmd.GenerateLoginCommands(evaClient),
		},
		cli.Command{
			Name:        "applications",
			Subcommands: cmd.GenerateApplicationCommands(evaClient),
		},
		cli.Command{
			Name:        "blob",
			Subcommands: cmd.GenerateBlobCommands(evaClient),
		},
		cli.Command{
			Name:        "products",
			Subcommands: cmd.GenerateProductCommands(evaClient),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%+v", err)
		fmt.Println()
	}
}

func getPathToHost() (string, error) {
	u, err := user.Current()

	if err != nil {
		return "", err
	}

	pathToDir := filepath.Join(u.HomeDir, ".eva-cli")

	if err := os.MkdirAll(pathToDir, os.ModePerm); err != nil {
		return "", err
	}

	pathToHost := filepath.Join(pathToDir, ".host")

	return pathToHost, nil
}
