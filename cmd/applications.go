package cmd

import (
	"fmt"

	"github.com/new-black/eva-cli/client"
	"github.com/urfave/cli"
)

func GenerateApplicationCommands(client client.Client) []cli.Command {
	return []cli.Command{
		cli.Command{
			Name: "list",
			Action: func(c *cli.Context) error {
				applications, err := client.GetApplications()

				if err != nil {
					return err
				}

				fmt.Printf("ID\tName")
				fmt.Println()

				for _, app := range applications.Result {
					fmt.Printf("%d\t%s", app.ID, app.Name)
					fmt.Println()
				}

				return nil
			},
		},
	}
}
