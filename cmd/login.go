package cmd

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/new-black/eva-cli/client"
	"github.com/urfave/cli"
)

func GenerateLoginCommands(client client.Client) []cli.Command {
	return []cli.Command{
		cli.Command{
			Name: "me",
			Action: func(c *cli.Context) error {
				u, err := client.GetCurrentUser()

				if err == nil {
					fmt.Printf("logged in as %+v", u)
					fmt.Println()
				}

				return err
			},
		},
		cli.Command{
			Name: "login",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "username, u"},
				cli.IntFlag{Name: "application, app"},
				cli.IntFlag{Name: "organization, ou"},
			},
			Action: func(c *cli.Context) error {
				password, err := terminal.ReadPassword(int(syscall.Stdin))

				if err != nil {
					return err
				}

				res, err := client.Login(c.String("username"), string(password), c.Int("ou"), c.Int("app"))

				if err != nil {
					return err
				}

				if res.Authentication != 2 /* Success */ {
					return fmt.Errorf("login failed, authentication result: %d", res.Authentication)
				}

				if err := client.StoreToken(res.AuthenticationToken); err != nil {
					return err
				}

				return nil
			},
		},
	}
}
